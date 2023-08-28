package http

import (
	"net/http"

	"github.com/KadirbekSharau/apprentice-backend/internal/auth"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	useCase auth.UseCase
}

func NewHandler(useCase auth.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

func (h *Handler) SignUpApplicant(c *gin.Context) {
	inp := new(auth.SignUpInput)

	if err := c.BindJSON(inp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := h.useCase.SignUpApplicant(c.Request.Context(), inp); err != nil {
		handleSignUpError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"Success": "Applicant User Registered Successfully"})
}

func (h *Handler) SignUpEmployer(c *gin.Context) {
	inp := new(auth.SignUpInput)

	if err := c.BindJSON(inp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := h.useCase.SignUpEmployer(c.Request.Context(), inp); err != nil {
		handleSignUpError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"Success": "Employer User Registered Successfully"})
}

func (h *Handler) SignIn(c *gin.Context) {
	inp := new(auth.SignInInput)

	if err := c.BindJSON(inp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	token, err := h.useCase.SignIn(c.Request.Context(), inp)
	if err != nil {
		if err == auth.ErrUserNotFound {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, auth.SignInResponse{Token: token})
}

// Handling HTTP status error
func handleSignUpError(c *gin.Context, err error) {
	switch err {
	case auth.ErrEmailAlreadyExists:
		c.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
	case auth.ErrInvalidEmailFormat:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email format"})
	case auth.ErrInvalidPassword:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid password"})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	}
}