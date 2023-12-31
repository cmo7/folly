package generics

import (
	"fmt"
	"folly/src/database"
	"folly/src/lib/common"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository[Entity common.Entity, DTO common.DTO] interface {
	Create(payload Entity) (Entity, error)
	Update(payload Entity) (Entity, error)
	Delete(payload Entity) error
	FindOne(id uuid.UUID, conditions common.SQLConditions, relations []string) (Entity, error)
	FindAll(pageable common.Pageable, conditions common.SQLConditions, relations []string, orderBys common.OrderBys) (*common.Page[Entity], error)
	FindOneRandom() (Entity, error)
	Exists(id uuid.UUID) (bool, error)
	Count(conditions common.SQLConditions) (int64, error)
}

type GenericRepository[Entity common.Entity, DTO common.DTO] struct{}

func NewGenericRepositoryGORM[Entity common.Entity, DTO common.DTO]() GenericRepository[Entity, DTO] {
	var service GenericRepository[Entity, DTO]
	return service
}

func (imp GenericRepository[Entity, DTO]) Create(payload Entity) (Entity, error) {
	err := database.DB.Create(&payload).Error
	return payload, err
}

func (imp GenericRepository[Entity, DTO]) Update(payload Entity) (Entity, error) {
	if payload.GetID() == uuid.Nil {
		return payload, fmt.Errorf("ID cannot be nil")
	}
	err := database.DB.
		Session(&gorm.Session{FullSaveAssociations: true}).
		Save(&payload).Error
	return payload, err
}

func (imp GenericRepository[Entity, DTO]) Delete(payload Entity) error {
	return database.DB.Delete(&payload).Error
}

func (imp GenericRepository[Entity, DTO]) FindOne(id uuid.UUID, conditions common.SQLConditions, relations []string) (Entity, error) {
	var entity Entity
	err := database.DB.
		Scopes(Preload(relations), Filter(conditions)).
		First(&entity, "id = ?", id).Error
	return entity, err
}

func (imp GenericRepository[Entity, DTO]) FindOneRandom() (Entity, error) {
	var entity Entity
	err := database.DB.Order("RANDOM()").First(&entity).Error
	return entity, err
}

func (imp GenericRepository[Entity, DTO]) FindAll(pageable common.Pageable, conditions common.SQLConditions, relations []string, orderBys common.OrderBys) (*common.Page[Entity], error) {
	var entities []Entity
	limit := pageable.Size
	offset := (pageable.Page - 1) * pageable.Size

	var count int64
	database.DB.Model(&entities).Count(&count)
	result := database.DB.
		Limit(limit).
		Offset(offset).
		Scopes(
			Preload(relations),
			Filter(conditions),
			Order(orderBys),
		).
		Find(&entities)
	var filtered int64
	database.DB.Model(&entities).
		Scopes(Filter(conditions)).
		Count(&filtered)

	return &common.Page[Entity]{
		Items:    entities,
		Page:     pageable.Page,
		Size:     pageable.Size,
		Total:    count,
		Filtered: filtered,
	}, result.Error
}

func (imp GenericRepository[Entity, DTO]) Exists(id uuid.UUID) (bool, error) {
	var entity Entity
	err := database.DB.First(&entity, "id = ?", id).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (imp GenericRepository[Entity, DTO]) Count(conditions common.SQLConditions) (int64, error) {
	var count int64
	var entity Entity
	err := database.DB.
		Model(entity).
		Scopes(Filter(conditions)).
		Count(&count).Error
	return count, err
}

func Filter(conditions common.SQLConditions) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		for _, condition := range conditions {
			if condition.Type == common.Or {
				db = db.Or(condition.Field+" "+string(condition.Comparator)+" ?", condition.Value)
				continue
			}
			db = db.Where(condition.Field+" "+string(condition.Comparator)+" ?", condition.Value)
		}
		return db
	}
}

func Order(orderBys common.OrderBys) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		for _, orderBy := range orderBys {
			if orderBy.Direction == common.Asc {
				db = db.Order(orderBy.Field + " asc")
			} else {
				db = db.Order(orderBy.Field + " desc")
			}
		}
		return db
	}
}

func Preload(relations []string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		for _, relation := range relations {
			db = db.Preload(relation)
		}
		return db
	}
}
