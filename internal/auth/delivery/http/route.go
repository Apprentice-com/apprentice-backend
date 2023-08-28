package http

import (
	"github.com/KadirbekSharau/apprentice-backend/internal/auth"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.Engine, uc auth.UseCase) {
	h := NewHandler(uc)
	authEndpoints := router.Group("/auth")
	{
		authEndpoints.POST("/sign-up/applicant", h.SignUpApplicant)
		authEndpoints.POST("/sign-up/employer", h.SignUpEmployer)
		authEndpoints.POST("/sign-in", h.SignIn)
	}
}