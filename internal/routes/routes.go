package routes

import (
	"github.com/gabrielagui373/obiwanapp-api/internal/config"
	"github.com/gabrielagui373/obiwanapp-api/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RouterDependecies struct {
	DB        *gorm.DB
	JWTConfig *config.JWTConfig
}

func SetupRouter(deps RouterDependecies) *gin.Engine {
	router := gin.Default()

	// Configuração de serviços
	authService := services.NewAuthService(deps.DB, deps.JWTConfig)

	// Configuração de rotas públicas
	setupPublicRoutes(router, authService)

	// Configuração de rotas protegidas
	setupProtectedRoutes(router, deps.DB, authService)

	return router
}
