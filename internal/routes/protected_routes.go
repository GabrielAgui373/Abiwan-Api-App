package routes

import (
	"github.com/gabrielagui373/obiwanapp-api/internal/middlewares"
	"github.com/gabrielagui373/obiwanapp-api/internal/repositories"
	"github.com/gabrielagui373/obiwanapp-api/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func setupProtectedRoutes(router *gin.Engine, db *gorm.DB, authService *services.AuthService) {
	protected := router.Group("/")
	protected.Use(middlewares.AuthMiddleware(authService))

	subjectRepo := repositories.NewSubjectRepository(db)
	topicRepo := repositories.NewTopicRepository(db)

	protectedRoutes := []ProtectedRouteGroup{
		NewSubjectRoutes(subjectRepo),
		NewTopicRoutes(topicRepo),
	}

	for _, routeGroup := range protectedRoutes {
		routeGroup.SetupRoutes(protected)
	}
}
