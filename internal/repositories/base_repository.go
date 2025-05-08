package repositories

import (
	"errors"
	"reflect"

	"gorm.io/gorm"
)

type Repository[T any] interface {
	FindAll(userID string) ([]T, error)
	FindByID(userID string, id string) (*T, error)
	Create(userID string, item *T) error
	Update(userID string, id string, item *T) error
	Delete(userID string, id string) error
	FindByFilters(userID string, filters map[string]interface{}) ([]T, error)
}

type BaseRepository[T any] struct {
	db *gorm.DB
}

func NewBaseRepository[T any](db *gorm.DB) *BaseRepository[T] {
	return &BaseRepository[T]{db: db}
}

func (r *BaseRepository[T]) FindAll(userID string) ([]T, error) {
	var items []T
	if err := r.db.Where("user_id = ?", userID).Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (r *BaseRepository[T]) FindByID(userID string, id string) (*T, error) {
	var item T
	if err := r.db.Where("id = ? AND user_id = ?", id, userID).First(&item).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *BaseRepository[T]) Create(userID string, item *T) error {
	// Usamos reflection para setar o UserID se o campo existir
	if err := setUserID(item, userID); err != nil {
		return err
	}
	return r.db.Create(item).Error
}

func (r *BaseRepository[T]) Update(userID string, id string, item *T) error {
	var existing T
	if err := r.db.Where("id = ? AND user_id = ?", id, userID).First(&existing).Error; err != nil {
		return err
	}

	// Atualiza o item garantindo que o user_id não seja alterado
	if err := setUserID(item, userID); err != nil {
		return err
	}

	result := r.db.Model(&item).Where("id = ? AND user_id = ?", id, userID).Updates(item)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *BaseRepository[T]) Delete(userID string, id string) error {
	result := r.db.Where("id = ? AND user_id = ?", id, userID).Delete(new(T))
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *BaseRepository[T]) FindByFilters(userID string, filters map[string]interface{}) ([]T, error) {
	var items []T
	query := r.db.Where("user_id = ?", userID)
	for key, value := range filters {
		query = query.Where(key+" = ?", value)
	}
	if err := query.Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

// Função auxiliar para setar o UserID usando reflection
func setUserID(item interface{}, userID string) error {
	// Usamos reflection para verificar e setar o campo UserID
	// Isso permite que o repositório seja genérico
	// Nota: Em produção, considere uma solução mais robusta ou gere código específico

	v := reflect.ValueOf(item).Elem()
	if v.Kind() != reflect.Struct {
		return errors.New("item must be a struct pointer")
	}

	field := v.FieldByName("UserID")
	if !field.IsValid() || !field.CanSet() {
		return errors.New("struct must have a settable UserID field")
	}

	if field.Kind() != reflect.String {
		return errors.New("UserID field must be of type uint")
	}

	field.Set(reflect.ValueOf(userID))
	return nil
}
