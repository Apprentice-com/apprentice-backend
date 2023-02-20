package routes

import (
	"github.com/KadirbekSharau/apprentice-backend/src/controllers"
	"github.com/KadirbekSharau/apprentice-backend/src/models"
	"github.com/KadirbekSharau/apprentice-backend/src/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

/* @description All Auth routes */
func InitAuthRoutes(db *gorm.DB, route *gin.Engine) {
	userController := controllers.NewUserController(services.NewUserService(models.NewUserRepository(db)))

	groupRoute := route.Group("/api/v1/auth")
	groupRoute.POST("/user/login", userController.UserLogin)
	groupRoute.POST("/user/register", userController.ActiveUserSeekerRegister)
}
