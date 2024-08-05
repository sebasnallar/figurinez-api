package infrastructure

import (
	"figurinez-api/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(deps *Dependencies) *gin.Engine {
	router := gin.Default()

	router.Use(gin.Recovery())
	router.Use(middleware.SecurityHeaders())
	router.Use(middleware.RequestID())

	deps.Handlers.HealthHandler.RegisterRoutes(router)

	return router
}
