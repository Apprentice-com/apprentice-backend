package http

import (
	jobpost "github.com/KadirbekSharau/apprentice-backend/internal/job_post"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.Engine, uc jobpost.UseCase) {
	h := NewHandler(uc)
	authEndpoints := router.Group("/job-post")
	{
		authEndpoints.POST("", h.CreateJobPost)
	}
}
