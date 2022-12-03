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

	fields, err := h.service.GetAllJobPosts()

	switch err {

	case "RESULTS_NOT_FOUND_404":
		util.APIResponse(ctx, "Data do not exist", http.StatusNotFound, http.MethodGet, nil)

	default:
		util.APIResponse(ctx, "Data found successfully", http.StatusOK, http.MethodGet, fields)
	}
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
	_, errAccount := h.service.CreateJobPost(&input)

	switch errAccount {
	case "CREATE_CONFLICT_409":
		util.APIResponse(ctx, "Already exist", http.StatusConflict, http.MethodPost, nil)
		return
	case "CREATE_FAILED_403":
		util.APIResponse(ctx, "Create new instance failed", http.StatusForbidden, http.MethodPost, nil)
		return
	default:
		util.APIResponse(ctx, "Create new instance successfully", http.StatusCreated, http.MethodPost, nil)
	}
}