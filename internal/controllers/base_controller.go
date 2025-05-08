package controllers

import (
	"net/http"

	"github.com/gabrielagui373/obiwanapp-api/internal/services"
	"github.com/gin-gonic/gin"
)

type BaseController[T any] struct {
	service services.Service[T]
}

func NewBaseController[T any](service services.Service[T]) *BaseController[T] {
	return &BaseController[T]{service: service}
}

func (bc *BaseController[T]) parseID(c *gin.Context) (string, bool) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return "", false
	}
	return id, true
}

func (bc *BaseController[T]) GetAll(c *gin.Context) {
	items, err := bc.service.GetAllItems(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, items)
}

func (bc *BaseController[T]) GetByID(c *gin.Context) {
	id, ok := bc.parseID(c)
	if !ok {
		return
	}

	item, err := bc.service.GetItemByID(id, c)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}
	c.JSON(http.StatusOK, item)
}

func (bc *BaseController[T]) Create(c *gin.Context) {
	var item T
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := bc.service.CreateItem(&item, c); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, item)
}

func (bc *BaseController[T]) Update(c *gin.Context) {
	id, ok := bc.parseID(c)
	if !ok {
		return
	}

	var item T
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := bc.service.UpdateItem(id, &item, c); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, item)
}

func (bc *BaseController[T]) Delete(c *gin.Context) {
	id, ok := bc.parseID(c)
	if !ok {
		return
	}

	if err := bc.service.DeleteItem(id, c); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func (bc *BaseController[T]) GetByFilters(c *gin.Context) {
	filters := make(map[string]interface{})
	for key, values := range c.Request.URL.Query() {
		filters[key] = values[0]
	}

	items, err := bc.service.GetItemsByFilters(filters, c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, items)
}
