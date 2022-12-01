package handlers

import (
	"net/http"

	"github.com/KadirbekSharau/apprentice-backend/src/dto"
	skillSetService "github.com/KadirbekSharau/apprentice-backend/src/services/skill-set"
	"github.com/KadirbekSharau/apprentice-backend/src/util"
	"github.com/gin-gonic/gin"
)

type SkillSetHandler interface {
	CreateSkillSet(ctx *gin.Context)
}

type skillSetHandler struct {
	service skillSetService.Service
}

func NewSkillSetHandler(service skillSetService.Service) *skillSetHandler {
	return &skillSetHandler{service: service}
}

/* Create Education Details Handler */
func (h *skillSetHandler) CreateSkillSet(ctx *gin.Context) {
	var input dto.CreateSkillSet
	ctx.ShouldBindJSON(&input)

	config := util.ErrorConfig{
		Options: map[string]util.ErrorMetaConfig{
			"Name required": {
				Tag:     "required",
				Field:   "name",
				Message: "institution name is required on body",
			},
		},
	}

	errResponse, errCount := util.GoValidator(&input, config.Options)

	if errCount > 0 {
		util.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodPost, errResponse)
		return
	}

	_, errAccount := h.service.CreateSkillSet(&input)

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