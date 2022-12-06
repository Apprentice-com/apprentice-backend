package routes

import (
	"github.com/KadirbekSharau/apprentice-backend/src/handlers"
	"github.com/KadirbekSharau/apprentice-backend/src/middleware"
	"github.com/KadirbekSharau/apprentice-backend/src/services/job-post"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

/* @description All Profile routes */
func InitJobPostRoutes(db *gorm.DB, route *gin.Engine) {
	var (
		repository = jobPostService.NewRepository(db)
		service    = jobPostService.NewService(repository)
		handler    = handlers.NewJobPostHandler(service)
	)

	groupRoute := route.Group("/api/v1/jobpost")
	groupRoute.GET("/", handler.GetAllJobPosts)
	groupRoute.GET("/:id", handler.GetJobPostByID)
	groupRoute.POST("/", middleware.Auth([]int{0, 2}), handler.CreateJobPost)
}