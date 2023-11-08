# FOLLY (Framework para aplicaciones web en go)

## Introducción

Folly es un framework para el desarrollo de aplicaciones web en go. Está basado en el framework [Fiber](https://gofiber.io/) y utiliza [Gorm](https://gorm.io/) como ORM.

La filosofía del framework es convención sobre configuración. Esto significa que el framework provee una serie de herramientas que permiten desarrollar aplicaciones web de forma rápida y sencilla.

Para ello se provee de una serie de fabricas que permiten crear controladores, repositorios y rutas genéricas para los modelos de la aplicación. Estos controladores, repositorios y rutas genéricas proveen las operaciones CRUD básicas para los modelos de la aplicación.

En caso de ser necesario, se pueden crear controladores, repositorios y rutas personalizadas para los modelos de la aplicación. Además se pueden extender las rutas genéricas para agregar funcionalidad.

## Instalación

## Uso

### CLI

El framework provee un CLI para facilitar la creación de modelos, controladores, repositorios y rutas. Para utilizar el CLI se debe ejecutar el comando `go run .` desde la carpeta raíz del proyecto.

El CLI provee su propia ayuda, para acceder a ella se debe ejecutar el comando `go run . help`.

### Modelos

Los modelos se componen al menos de una Entidad y un DTO (Data Transfer Object). La Entidad es la representación de la tabla en la base de datos y el DTO es la representación de la entidad en la API.

#### Entidad

Las Entidades son structs que representan una tabla en la base de datos. Para que una struct sea considerada una Entidad debe cumplir con los siguientes requisitos:

- Debe ser una struct
- Debe implementar la interfaz `common.Entity`
  - Esta interfaz contiene los métodos `GetId()`, `SetId()` y `ToDto()`
  - Debe tener un campo `Id` de tipo `uuid.UUID`

Para facilitar las tareas de creación de Entidades, se recomienda incrustar el struct `common.CommonEntity`.

Se utiliza el ORM Gorm, por lo que se pueden utilizar todas sus etiquetas para configurar la Entidad. Para más información sobre las etiquetas de Gorm, consultar [la documentación](https://gorm.io/docs/models.html).

A diferencia de lo que figura en la documentación de Gorm usamos campos `uuid.UUID` en lugar de `uint` para el campo `ID`. Todos los campos `ID` deben ser de tipo `uuid.UUID`.

Ejemplo:

```go
type Role struct {
    common.CommonEntity `gorm:"embedded"`
    Name                string `gorm:"unique;not null"`
    Users               []User `gorm:"many2many:user_roles;"`
}
```

#### DTO

Los DTOs son structs que representan una Entidad en la API. Para que una struct sea considerada un DTO debe cumplir con los siguientes requisitos:

- Debe ser una struct
- Debe implementar la interfaz `common.Dto`
  - Esta interfaz contiene los métodos `GetId()`, `SetId()` y `ToEntity()`
  - Debe tener un campo `Id` de tipo `uuid.UUID`

Los campos de un DTO deben ser públicos para que puedan ser serializados por el paquete `encoding/json`.

Es recomendable que los campos de un DTO sean del mismo tipo que los campos de la Entidad que representa.

Es recomendable que los campos estén etiquetados con `json:"{nombre}"` para que el serializador pueda identificarlos correctamente. Los nombres de los campos deberían ser los mismos que los de la Entidad que representa, pero siguiento la convención de nombres de JSON (camelCase).

Campos que no se deseen serializar pueden ser etiquetados con `json:"-"` aunque es recomendable que estos campos no existan en el DTO, por ejemplo passwords y otros campos sensibles.

Para facilitar las tareas de creación de DTOs, se recomienda incrustar el struct `common.CommonDto`.

Ejemplo:

```go
type RoleDTO struct {
    common.CommonDTO `json:",inline,omitempty"`
    Name             string    `json:"name"`
    Users            []UserDTO `json:"users,omitempty"`
}
```

### Migraciones

Los modelos se migran automáticamente utilizando el ORM Gorm. Para ello se debe llamar a la función `database.Migrate`. El CLI provee un comando para migrar los modelos de la aplicación.

Los modelos han de ser registrados como tareas de migración. Se recomienda utilizar una función init() para ello:

TODO: Mejorar la documentación de migraciones

```go
func init() {
	database.RegisterMigration(&database.MigrationTask{
		Model:           &User{},
		DropOnFlush:     true,
		TruncateOnFlush: true,
	})
}
```

### Repositorios

Los repositorios son structs que implementan la interfaz `generics.Repository`. Esta interfaz contiene los métodos básicos para interactuar con la base de datos.

Los repositorios genéricos proveen los siguientes métodos:

TODO: Mejorar la documentación de repositorios

| Nombre | Descripción | Parámetros | Retorno |
| -- | ------ | ---- | ----------- |
| FindAll | Retorna una lista de entidades |  |  |
| FindOne | Retorna una entidad específica |  |  |
| Create | Crea una nueva entidad |  |  |
| Update | Actualiza una entidad específica |  |  |
| Delete | Elimina una entidad específica |  |  |
| FindOneRandom | Retorna una entidad aleatoria |  |  |
| Exists | Retorna true si existe una entidad con el id especificado |  |  |
| Count | Retorna la cantidad de entidades que cumplen una serie de condiciones |  |  |

Para crear un repositorio genérico y opcionalmente extenderlo, se debe crear un nuevo interface que incruste `generics.Repository[Entity, DTO]` y luego crear un struct que implemente ese interface.

Ejemplo sin métodos adicionales:

```go
type RoleRepository struct {
	generics.GenericRepository[*models.Role, *models.RoleDTO]
}

var RoleRepositoryGORM RoleRepository = RoleRepository{}
```

Ejemplo con métodos adicionales:

```go
type UserRepository struct {
	generics.GenericRepository[*models.User, *models.UserDTO]
	// Add custom method signatures here
	FindByEmail  func(email string, relations []string) (models.User, error)
	IsEmailTaken func(email string) bool
}

var UserRepositoryGORM UserRepository = UserRepository{
	// Add custom method imlementations here
	FindByEmail: func(email string, relations []string) (models.User, error) {
		var user models.User
		result := database.DB.
			Scopes(generics.Preload(relations)).
			Where("email = ?", email).
			First(&user)
		return user, result.Error
	},
	IsEmailTaken: func(email string) bool {
		var count int64
		database.DB.Model(&models.User{}).Where("email = ?", email).Count(&count)
		return count > 0
	},
}
```

### Rutas y controladores

El controlador y rutas genéricos exponen los siguientes endpoints:

| Nombre | Método | Ruta | Descripción |
| -- | ------ | ---- | ----------- |
| GetAll | GET | {host}/api/v1/{recurso}/ | Retorna una lista de recursos |
| Get | GET | {host}/api/v1/{recurso}/{id} | Retorna un recurso específico |
| Create | POST | {host}/api/v1/{recurso}/ | Crea un nuevo recurso |
| Update | PUT | {host}/api/v1/{recurso}/{id} | Actualiza un recurso específico |
| Delete | DELETE | {host}/api/v1/{recurso}/{id} | Elimina un recurso específico |

El método GetAll pueden recibir los siguientes parámetros de consulta:

| Nombre | Tipo | Descripción | Ejemplo |
| -- | ------ | ---- | ----------- |
| page | int | Número de página | `page=1` |
| size | int | Tamaño de página | `size=10` |
| orders | string | Campo por el que se ordenará la lista | `orders=name:asc,apellido` |
| filters | string | Filtros a aplicar | `filters=name:eq:Juan,apellido:like:Gonzalez%` |
| relations | string | Relaciones a incluir (pueden ser anidadas con .) | `relations=roles,roles.permisos` |

Algunas urls de ejemplo:

#### Descripción

> Retornará la 3ª página de usuarios, ordenados por campo nombre en orden descendente, incluyendo los roles de cada usuario, filtrando solo aquellos cuyo nombre empiece con "Juan" y con un tamaño de página de 20 elementos.

#### URI

> `GET http://localhost:3000/api/v1/users?orders=first-name:desc?relations=roles?filters=first-name:like:Juan%?page=3?size=20`


#### Sintaxis Detallada de los parámetros de consulta

##### orders

> `orders={campo}:{asc|desc}`

> `orders={campo1}:{asc|desc},{campo2}:{asc|desc},...`

##### filters

> `filters={campo}:{eq|neq|gt|gte|lt|lte|like|nlike}:{valor}`

> `filters={campo1}:{eq|neq|gt|gte|lt|lte|like|nlike}:{valor1},{and|or}:{campo2}:{eq|neq|gt|gte|lt|lte|like|nlike}:{valor2},...`

##### relations

> `relations={campo1},{campo2},...`

> `relations={campo1}.{campoAnidado1},...`

> `relations={campo1}.{campoAnidado1},{campo2}.{campoAnidado2},...`

#### Controladores

Para crear un controlador genérico y registrar sus rutas se puede usar la función `RegisterController` de la siguiente manera:

```go
	RegisterController(generics.NewController[*models.User, *models.UserDTO](generics.ResourceNames{
		Singular: "user",
		Plural:   "users",
	}))
```
`generics.NewController` recibe los parámetros de tipo que representan la Entidad y el DTO, y recibe un struct `generics.ResourceNames` que contiene los nombres del recurso en singular y plural.

De forma transparente `NewController` crea un nuevo repositorio genérico y lo asigna al controlador.

`RegisterController` registra las rutas del controlador en el router de la aplicación.

## TODO

### Documentación

- [ ] Agregar documentación de la CLI
- [ ] Agregar documentación de la configuración
- [ ] Agregar documentación de los middlewares
- [ ] Agregar documentación de los tests
- [ ] Agregar documentación de los logs
- [ ] Agregar documentación de los errores

### Funcionalidad

- [ ] Agregar soporte para autorización
- [ ] Mejorar soporte para autenticación
- [ ] Mejorar interfaz del paquete `lib` para que sea más fácil de usar
- [ ] Agregar soporte para subir archivos
- [ ] Agregar soporte para internacionalización
- [x] Agregar soporte para paginación
- [x] Agregar soporte para ordenamiento
- [x] Agregar soporte para filtros
- [x] Agregar soporte para relaciones
- [x] Agregar soporte para migraciones
- [x] Agregar soporte para logs
