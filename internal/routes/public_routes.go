package routes

import (
	"github.com/gabrielagui373/obiwanapp-api/internal/controllers"
	"github.com/gabrielagui373/obiwanapp-api/internal/services"
	"github.com/gin-gonic/gin"
)

type AuthRoutes struct {
	authService *services.AuthService
}

func NewAuthRoutes(authService *services.AuthService) *AuthRoutes {
	return &AuthRoutes{authService: authService}
}

func (r *AuthRoutes) SetupRoutes(router *gin.Engine) {
	authController := controllers.NewAuthController(r.authService)

	public := router.Group("/auth")
	{
		public.POST("/register", authController.Register)
		public.POST("/login", authController.Login)
		public.POST("/refresh", authController.RefreshToken)
	}
}

func setupPublicRoutes(router *gin.Engine, authService *services.AuthService) {
	publicRoutes := []PublicRouteGroup{
		NewAuthRoutes(authService),
		//outros grupos de rotas públicas aqui
	}

	for _, routeGroup := range publicRoutes {
		routeGroup.SetupRoutes(router)
	}
}
