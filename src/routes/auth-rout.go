package routes

import (
	"github.com/KadirbekSharau/apprentice-backend/src/services/auth"
	"github.com/KadirbekSharau/apprentice-backend/src/handlers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

/* @description All Auth routes */
func InitAuthRoutes(db *gorm.DB, route *gin.Engine) {
	var (
		repository = authService.NewRepository(db)
		service    = authService.NewService(repository)
		authHandler    = handlers.NewAuthHandler(service)
	)

	groupRoute := route.Group("/api/v1/auth")
	groupRoute.POST("/user/login", authHandler.UserLogin)
	groupRoute.POST("/user/register", authHandler.ActiveUserSeekerRegister)
}
