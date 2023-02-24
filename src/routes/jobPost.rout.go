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
func InitJobPostRoutes(db *gorm.DB, route *gin.Engine) {
	controller := controllers.NewJobPostController(services.NewJobPostService(models.NewJobPostRepository(db)))

	groupRoute := route.Group("/api/v1/jobpost")
	groupRoute.GET("/", controller.GetAllJobPosts)
	groupRoute.GET("/:id", controller.GetJobPostByID)
	groupRoute.POST("/", middleware.Auth([]int{0, 2}), controller.CreateJobPost)
}