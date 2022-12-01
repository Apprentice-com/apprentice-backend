package handlers

import (
	"net/http"

	"github.com/KadirbekSharau/apprentice-backend/src/dto"
	profileService "github.com/KadirbekSharau/apprentice-backend/src/services/profile"
	"github.com/KadirbekSharau/apprentice-backend/src/util"
	"github.com/gin-gonic/gin"
)

type ProfileHandler interface {
	GetSeekerProfile(ctx *gin.Context)
	CreateEducationDetails(ctx *gin.Context)
}

type profileHandler struct {
	service profileService.Service
}

func NewProfileHandler(service profileService.Service) *profileHandler {
	return &profileHandler{service: service}
}

/* Get Seeker Profile Handler */
func (h *profileHandler) GetSeekerProfile(ctx *gin.Context) {

	var input dto.GetSeekerProfile
	ctx.Params.ByName("user_id")

	config := util.ErrorConfig{
		Options: map[string]util.ErrorMetaConfig{
			"ID required": {
				Tag:     "required",
				Field:   "ID",
				Message: "id is required on param",
			},
		},
	}

	errResponse, errCount := util.GoValidator(&input, config.Options)

	if errCount > 0 {
		util.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodGet, errResponse)
		return
	}

	result, err := h.service.GetSeekerProfile(&input)

	switch err {

	case "RESULT_STUDENT_NOT_FOUND_404":
		util.APIResponse(ctx, "data is not exist or deleted", http.StatusNotFound, http.MethodGet, nil)
		return

	default:
		util.APIResponse(ctx, "Result data successfully", http.StatusOK, http.MethodGet, result)
	}
}

/* Create Education Details Handler */
func (h *profileHandler) CreateEducationDetails(ctx *gin.Context) {
	var input dto.CreateEducationDetails
	ctx.ShouldBindJSON(&input)

	config := util.ErrorConfig{
		Options: map[string]util.ErrorMetaConfig{
			"Institution name required": {
				Tag:     "required",
				Field:   "institution_name",
				Message: "institution name is required on body",
			},
			"User ID required": {
				Tag:     "required",
				Field:   "user_id",
				Message: "user id is required on body",
			},
		},
	}

	errResponse, errCount := util.GoValidator(&input, config.Options)

	if errCount > 0 {
		util.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodPost, errResponse)
		return
	}

	_, errAccount := h.service.CreateEducationDetails(&input)

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
