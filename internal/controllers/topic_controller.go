package controllers

import (
	"net/http"

	"github.com/gabrielagui373/obiwanapp-api/internal/models"
	"github.com/gabrielagui373/obiwanapp-api/internal/services"
	"github.com/gin-gonic/gin"
)

type TopicController struct {
	*BaseController[models.Topic]
	service *services.TopicService
}

func NewTopicController(service *services.TopicService) *TopicController {
	return &TopicController{
		BaseController: NewBaseController(service),
		service:        service,
	}
}

func (tc *TopicController) GetAll(c *gin.Context) {
	items, err := tc.service.GetAllItems(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, items)
}

func (tc *TopicController) GetByID(c *gin.Context) {
	id, ok := tc.parseID(c)
	if !ok {
		return
	}

	item, err := tc.service.GetItemByID(id, c)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}
	c.JSON(http.StatusOK, item)
}
