package http

import (
	"net/http"

	middleware "github.com/KadirbekSharau/apprentice-backend/internal/auth/delivery/http"
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
	err := h.useCase.CreateJobPostByEmployer(c.Request.Context(), &input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Job Post"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "New Job Post created successfully"})
}

func (h *Handler) GetAllJobPostsByEmployerID(c *gin.Context) {
	employerID := middleware.GetUserIdFromContext(c) // Get employer ID from the context

    jobPosts, err := h.useCase.GetAllJobPostsByEmployerID(c, employerID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
        return
    }

    c.JSON(http.StatusOK, jobPosts)
}

func (h *Handler) GetAllJobPosts(c *gin.Context) {
    jobPosts, err := h.useCase.GetAllJobPosts(c)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
        return
    }

    c.JSON(http.StatusOK, jobPosts)
}