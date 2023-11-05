package generics

import (
	"folly/src/database"
	"folly/src/lib/common"
	"folly/src/lib/helpers"
)

type AttributeMap map[string]interface{}

type GeneratorMap map[string]func() interface{}

type Factory[T common.Entity] interface {
	// MakeOne creates a single instance of the model and returns it
	MakeOne() (*T, error)
	// MakeMany creates n instances of the model and returns them
	MakeMany(n int) (*[]T, error)
	// CreateOne creates a single instance of the model and saves it to the database
	CreateOne() (*T, error)
	// CreateMany creates n instances of the model and saves them to the database
	CreateMany(n int) (*[]T, error)
	// MakeOneWith creates a single instance of the model and returns it
	// with the provided attributes
	MakeOneWith(attributes AttributeMap) (*T, error)
	// MakeManyWith creates n instances of the model and returns them
	// with the provided attributes
	MakeManyWith(n int, attributes AttributeMap) (*[]T, error)
	// CreateOneWith creates a single instance of the model and saves it to the database
	// with the provided attributes
	CreateOneWith(attributes AttributeMap) (*T, error)
	// CreateManyWith creates n instances of the model and saves them to the database
	// with the provided attributes
	CreateManyWith(n int, attributes AttributeMap) ([]*T, error)
}

type GenericFactory[T common.Entity] struct {
	generators GeneratorMap
}

func NewFactory[T common.Entity](generators GeneratorMap) GenericFactory[T] {
	var factory GenericFactory[T]
	factory.generators = generators
	return factory
}

func (imp GenericFactory[T]) MakeOne() (*T, error) {
	defaultValues := map[string]interface{}{}
	for key, generator := range imp.generators {
		defaultValues[key] = generator()
	}
	return imp.MakeOneWith(defaultValues)
}

func (imp GenericFactory[T]) MakeMany(n int) (*[]T, error) {
	models := []T{}
	for i := 0; i < n; i++ {
		model, err := imp.MakeOne()
		if err != nil {
			return nil, err
		}
		models = append(models, *model)
	}
	return &models, nil
}

func (imp GenericFactory[T]) CreateOne() (*T, error) {
	model, err := imp.MakeOne()
	if err != nil {
		return nil, err
	}
	err = database.DB.Create(&model).Error
	return model, err
}

func (imp GenericFactory[T]) CreateMany(n int) (*[]T, error) {
	models := []T{}
	for i := 0; i < n; i++ {
		model, err := imp.MakeOne()
		if err != nil {
			return nil, err
		}
		models = append(models, *model)
	}
	err := database.DB.Create(&models).Error
	return &models, err
}

func (imp GenericFactory[T]) MakeOneWith(attributes AttributeMap) (*T, error) {
	entityData := map[string]interface{}{}

	for key, generator := range imp.generators {
		if _, ok := attributes[key]; !ok {
			entityData[key] = generator()
		} else {
			entityData[key] = attributes[key]
		}
	}
	var entity T
	err := helpers.PopulateStruct(entityData, &entity)
	return &entity, err
}

func (imp GenericFactory[T]) MakeManyWith(n int, attributes AttributeMap) ([]*T, error) {
	var entities []*T
	for i := 0; i < n; i++ {
		entity, err := imp.MakeOneWith(attributes)
		if err != nil {
			return nil, err
		}
		entities = append(entities, entity)
	}
	return entities, nil
}

func (imp GenericFactory[T]) CreateOneWith(attributes AttributeMap) (*T, error) {
	entity, err := imp.MakeOneWith(attributes)
	if err != nil {
		return nil, err
	}
	err = database.DB.Create(entity).Error
	return entity, err
}

func (imp GenericFactory[T]) CreateManyWith(n int, attributes AttributeMap) ([]*T, error) {
	entities := []*T{}
	for i := 0; i < n; i++ {
		entity, err := imp.MakeOneWith(attributes)
		if err != nil {
			return nil, err
		}
		entities = append(entities, entity)
	}
	err := database.DB.Create(entities).Error
	return entities, err
}
