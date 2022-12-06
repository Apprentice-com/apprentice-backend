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
	input.UserID = ctx.Params.ByName("user_id")
	config := util.ErrorConfig{
		Options: map[string]util.ErrorMetaConfig{
			"ID required": {
				Tag:     "required",
				Field:   "ID",
				Message: "id is required on param",
			},
		},
	}
	if errResponse, errCount := util.GoValidator(&input, config.Options); errCount > 0 {
		util.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodPost, errResponse)
		return
	}
	data, status, err := h.service.GetSeekerProfile(&input)
	util.APIResponse(ctx, err, status, http.MethodGet, data)
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
	if errResponse, errCount := util.GoValidator(&input, config.Options); errCount > 0 {
		util.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodPost, errResponse)
		return
	}
	data, status, err := h.service.CreateEducationDetails(&input)
	util.APIResponse(ctx, err, status, http.MethodGet, data)
}
