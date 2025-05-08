package routes

import (
	"github.com/gabrielagui373/obiwanapp-api/internal/controllers"
	"github.com/gabrielagui373/obiwanapp-api/internal/repositories"
	"github.com/gabrielagui373/obiwanapp-api/internal/services"
	"github.com/gin-gonic/gin"
)

type TopicRoutes struct {
	topicRepo *repositories.TopicRepository
}

func NewTopicRoutes(topicRepo *repositories.TopicRepository) *TopicRoutes {
	return &TopicRoutes{topicRepo: topicRepo}
}

func (r *TopicRoutes) SetupRoutes(router *gin.RouterGroup) {
	topicService := services.NewTopicService(r.topicRepo)
	topicController := controllers.NewTopicController(topicService)

	subjectGroup := router.Group("/topics")
	{
		subjectGroup.GET("/", topicController.GetAll)
		subjectGroup.GET("/:id", topicController.GetByID)
		subjectGroup.POST("/", topicController.Create)
		subjectGroup.PUT("/:id", topicController.Update)
		subjectGroup.DELETE("/:id", topicController.Delete)

	}
}
