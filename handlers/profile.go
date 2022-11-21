package handlers

import (
	"net/http"

	"github.com/KadirbekSharau/apprentice-backend/services/profile"
	"github.com/KadirbekSharau/apprentice-backend/util"
	"github.com/gin-gonic/gin"
)

type ProfileHandler interface {
	GetSeekerProfile(ctx *gin.Context)
}

type profileHandler struct {
	service profileService.Service
}

func NewProfileHandler(service profileService.Service) *profileHandler {
	return &profileHandler{service: service}
}

/* Get Seeker Profile Handler */
func (h *profileHandler) GetSeekerProfile(ctx *gin.Context) {

	var input profileService.InputGetSeekerProfile
	ctx.Params.ByName("user_id")

	config := util.ErrorConfig{
		Options: []util.ErrorMetaConfig{
			{
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