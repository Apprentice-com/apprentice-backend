package controllers

import (
	"net/http"

	"github.com/KadirbekSharau/apprentice-backend/src/dto"
	"github.com/KadirbekSharau/apprentice-backend/src/services"
	"github.com/KadirbekSharau/apprentice-backend/src/util"
	"github.com/gin-gonic/gin"
)

var config = util.ErrorConfig{
	Options: map[string]util.ErrorMetaConfig{
		"Institution name required": {
			Tag:     "required",
			Field:   "job_name",
			Message: "institution name is required on body",
		},
	},
}

type JobPostController interface {
	GetAllJobPosts(ctx *gin.Context)
	CreateJobPost(ctx *gin.Context)
}

type jobPostController struct {
	service *services.JobPostService
}

func NewJobPostController(service *services.JobPostService) *jobPostController {
	return &jobPostController{service: service}
}

/* Get All Job Posts Handler */
func (h *jobPostController) GetAllJobPosts(ctx *gin.Context) {
	data, status, err := h.service.GetAllJobPosts()
	util.APIResponse(ctx, err, status, http.MethodGet, data)
}

/* Create Education Details Handler */
func (h *jobPostController) CreateJobPost(ctx *gin.Context) {
	var input dto.CreateJobPost
	ctx.ShouldBindJSON(&input)
	if errResponse, errCount := util.GoValidator(&input, config.Options); errCount > 0 {
		util.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodPost, errResponse)
		return
	}
	data, status, err := h.service.CreateJobPost(&input)
	util.APIResponse(ctx, err, status, http.MethodGet, data)

}

/* Get Job Post By ID Handler */
func (h *jobPostController) GetJobPostByID(ctx *gin.Context) {
	data, status, err := h.service.GetJobPostByID(ctx.Params.ByName("id"))
	util.APIResponse(ctx, err, status, http.MethodGet, data)
}