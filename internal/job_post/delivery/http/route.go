package http

import (
	jobpost "github.com/KadirbekSharau/apprentice-backend/internal/job_post"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.Engine, uc jobpost.UseCase) {
	h := NewHandler(uc)
	jobPostEndpoints := router.Group("/job-post")
	{
		jobPostEndpoints.POST("", h.CreateJobPost)
		jobPostEndpoints.GET("", h.GetAllJobPosts)
		jobPostEndpoints.GET("/:id", h.GetAllJobPostsByEmployerID)
	}
}
