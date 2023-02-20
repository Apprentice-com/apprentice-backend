package controllers

import (
	"net/http"

	"github.com/KadirbekSharau/apprentice-backend/src/dto"
	"github.com/KadirbekSharau/apprentice-backend/src/services"
	"github.com/KadirbekSharau/apprentice-backend/src/util"
	"github.com/gin-gonic/gin"
)

type ProfileController interface {
	GetSeekerProfile(ctx *gin.Context)
	CreateEducationDetails(ctx *gin.Context)
}

type profileController struct {
	service *services.ProfileService
}

func NewProfileController(service *services.ProfileService) *profileController {
	return &profileController{service: service}
}

/* Get Seeker Profile Handler */
func (h *profileController) GetSeekerProfile(ctx *gin.Context) {
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