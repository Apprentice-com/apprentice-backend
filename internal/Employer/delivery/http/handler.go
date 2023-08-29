package http

import (
	"net/http"

	"github.com/KadirbekSharau/apprentice-backend/internal/employer"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	useCase employer.UseCase
}

func NewHandler(useCase employer.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

func (h *Handler) CreateEmployerProfile(c *gin.Context) {
	// Parse request body and validate input
	var input employer.CreateEmployerProfileInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Call use case method
	err := h.useCase.CreateEmployerProfile(c.Request.Context(), &input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create employer profile"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Employer profile created successfully"})
}
