package http

import (
	"net/http"

	jobpost "github.com/KadirbekSharau/apprentice-backend/internal/job_post"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	useCase jobpost.UseCase
}

func NewHandler(useCase jobpost.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

func (h *Handler) CreateJobPost(c *gin.Context) {
	var input jobpost.CreateJobPostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Call use case method
	err := h.useCase.CreateJobPost(c.Request.Context(), &input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Job Post"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "New Job Post created successfully"})
}
