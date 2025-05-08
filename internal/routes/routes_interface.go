package routes

import "github.com/gin-gonic/gin"

type PublicRouteGroup interface {
	SetupRoutes(*gin.Engine)
}

type ProtectedRouteGroup interface {
	SetupRoutes(*gin.RouterGroup)
}
