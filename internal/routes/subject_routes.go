package routes

import (
	"github.com/gabrielagui373/obiwanapp-api/internal/controllers"
	"github.com/gabrielagui373/obiwanapp-api/internal/repositories"
	"github.com/gabrielagui373/obiwanapp-api/internal/services"
	"github.com/gin-gonic/gin"
)

type SubjectRoutes struct {
	subjectRepo *repositories.SubjectRepository
}

func NewSubjectRoutes(subjectRepo *repositories.SubjectRepository) *SubjectRoutes {
	return &SubjectRoutes{subjectRepo: subjectRepo}
}

func (r *SubjectRoutes) SetupRoutes(router *gin.RouterGroup) {
	subjectService := services.NewSubjectService(r.subjectRepo)
	subjectController := controllers.NewSubjectController(subjectService)

	subjectGroup := router.Group("/subjects")
	{
		subjectGroup.GET("/", subjectController.GetAll)
		subjectGroup.GET("/:id", subjectController.GetByID)
		subjectGroup.POST("/", subjectController.Create)
		subjectGroup.PUT("/:id", subjectController.Update)
		subjectGroup.DELETE("/:id", subjectController.Delete)

	}
}
