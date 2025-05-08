package services

import (
	"fmt"

	"github.com/gabrielagui373/obiwanapp-api/internal/models"
	"github.com/gabrielagui373/obiwanapp-api/internal/repositories"
	"github.com/gin-gonic/gin"
)

type Service[T any] interface {
	GetAllItems(c *gin.Context) ([]T, error)
	GetItemByID(id string, c *gin.Context) (*T, error)
	CreateItem(item *T, c *gin.Context) error
	UpdateItem(id string, item *T, c *gin.Context) error
	DeleteItem(id string, c *gin.Context) error
	GetItemsByFilters(filters map[string]interface{}, c *gin.Context) ([]T, error)
}

type BaseService[T any] struct {
	repo repositories.Repository[T]
}

func NewBaseService[T any](repo repositories.Repository[T]) *BaseService[T] {
	return &BaseService[T]{repo: repo}
}

func (s *BaseService[T]) GetAllItems(c *gin.Context) ([]T, error) {
	user, err := s.getUser(c)
	if err != nil {
		return nil, err
	}
	return s.repo.FindAll(user.ID)
}

func (s *BaseService[T]) GetItemByID(id string, c *gin.Context) (*T, error) {
	user, err := s.getUser(c)
	if err != nil {
		return nil, err
	}
	return s.repo.FindByID(user.ID, id)
}

func (s *BaseService[T]) CreateItem(item *T, c *gin.Context) error {
	user, err := s.getUser(c)
	if err != nil {
		return err
	}
	return s.repo.Create(user.ID, item)
}

func (s *BaseService[T]) UpdateItem(id string, item *T, c *gin.Context) error {
	user, err := s.getUser(c)
	if err != nil {
		return err
	}
	return s.repo.Update(user.ID, id, item)
}

func (s *BaseService[T]) DeleteItem(id string, c *gin.Context) error {
	user, err := s.getUser(c)
	if err != nil {
		return err
	}
	return s.repo.Delete(user.ID, id)
}

func (s *BaseService[T]) GetItemsByFilters(filters map[string]interface{}, c *gin.Context) ([]T, error) {
	user, err := s.getUser(c)
	if err != nil {
		return nil, err
	}
	return s.repo.FindByFilters(user.ID, filters)
}

func (s *BaseService[T]) getUser(c *gin.Context) (*models.User, error) {
	user, exists := c.Get("user")
	if !exists {
		return nil, fmt.Errorf("user not found in context")
	}

	return user.(*models.User), nil
}
