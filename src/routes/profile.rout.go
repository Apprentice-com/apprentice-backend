package routes

import (
	"github.com/KadirbekSharau/apprentice-backend/src/controllers"
	"github.com/KadirbekSharau/apprentice-backend/src/middleware"
	"github.com/KadirbekSharau/apprentice-backend/src/models"
	"github.com/KadirbekSharau/apprentice-backend/src/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

/* @description All Profile routes */
func InitProfileRoutes(db *gorm.DB, route *gin.Engine) {
	controller := controllers.NewProfileController(services.NewProfileService(models.NewProfileRepository(db)))

	groupRoute := route.Group("/api/v1/")
	groupRoute.GET("/profile", middleware.Auth([]int{1}), controller.GetSeekerProfile)
	//groupRoute.POST("/educationdetails", middleware.Auth([]int{1}), handler.CreateEducationDetails)
}
