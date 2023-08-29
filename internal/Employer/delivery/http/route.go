package http

import (
	"github.com/KadirbekSharau/apprentice-backend/internal/employer"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.Engine, uc employer.UseCase) {
	h := NewHandler(uc)
	authEndpoints := router.Group("/profile")
	{
		authEndpoints.POST("/employer", h.CreateEmployerProfile)
	}
}
