package routes

import (
	"github.com/KadirbekSharau/apprentice-backend/src/handlers"
	"github.com/KadirbekSharau/apprentice-backend/src/middleware"
	profileService "github.com/KadirbekSharau/apprentice-backend/src/services/profile"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

/* @description All Profile routes */
func InitProfileRoutes(db *gorm.DB, route *gin.Engine) {
	var (
		repository = profileService.NewRepository(db)
		service    = profileService.NewService(repository)
		handler    = handlers.NewProfileHandler(service)
	)

	groupRoute := route.Group("/api/v1/")
	groupRoute.GET("/profile", middleware.Auth([]int{1}), handler.GetSeekerProfile)
	groupRoute.POST("/educationdetails" /*middleware.Auth([]int{1}),*/, handler.CreateEducationDetails)
}
