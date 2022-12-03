package handlers

import (
	"net/http"

	"github.com/KadirbekSharau/apprentice-backend/src/dto"
	"github.com/KadirbekSharau/apprentice-backend/src/services/job-post"
	"github.com/KadirbekSharau/apprentice-backend/src/util"
	"github.com/gin-gonic/gin"
)

type JobPostHandler interface {
	GetAllJobPosts(ctx *gin.Context)
	CreateJobPost(ctx *gin.Context)
}

type jobPostHandler struct {
	service jobPostService.Service
}

func NewJobPostHandler(service jobPostService.Service) *jobPostHandler {
	return &jobPostHandler{service: service}
}

/* Get All Job Posts Handler */
func (h *jobPostHandler) GetAllJobPosts(ctx *gin.Context) {
	data, status, err := h.service.GetAllJobPosts()
	util.APIResponse(ctx, err, status, http.MethodGet, data)
}

/* Create Education Details Handler */
func (h *jobPostHandler) CreateJobPost(ctx *gin.Context) {
	var input dto.CreateJobPost
	ctx.ShouldBindJSON(&input)

	config := util.ErrorConfig{
		Options: map[string]util.ErrorMetaConfig{
			"Institution name required": {
				Tag:     "required",
				Field:   "job_name",
				Message: "institution name is required on body",
			},
		},
	}
	errResponse, errCount := util.GoValidator(&input, config.Options)
	if errCount > 0 {
		util.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodPost, errResponse)
		return
	}
	data, status, err := h.service.CreateJobPost(&input)
	util.APIResponse(ctx, err, status, http.MethodGet, data)

}