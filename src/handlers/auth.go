package handlers

import (
	"net/http"

	"github.com/KadirbekSharau/apprentice-backend/src/handlers/helpers"
	"github.com/KadirbekSharau/apprentice-backend/src/services/auth"
	"github.com/KadirbekSharau/apprentice-backend/src/util"
	"github.com/KadirbekSharau/apprentice-backend/src/dto"
	"github.com/gin-gonic/gin"
)

type AuthHandler interface {
	UserLogin(ctx *gin.Context)
	ActiveUserSeekerRegister(ctx *gin.Context)
}

type authHandler struct {
	service authService.Service
}

func NewAuthHandler(service authService.Service) AuthHandler {
	return &authHandler{service: service}
}

/* User Login Handler */
func (h *authHandler) UserLogin(ctx *gin.Context) {
	var input dto.InputLogin
	ctx.ShouldBindJSON(&input)
	if errResponse, errCount := util.GoValidator(&input, helpers.AuthConfig.Options); errCount > 0 {
		util.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodPost, errResponse)
		return
	}
	data, status, err := h.service.UserLogin(&input)
	helpers.UserLoginTokenHandler(ctx, status, err, data)
}

/* Active User Register Handler */
func (h *authHandler) ActiveUserSeekerRegister(ctx *gin.Context) {
	var input dto.InputUserSeekerRegister
	ctx.ShouldBindJSON(&input)
	conf := helpers.AuthConfig.Options
	conf["Password minimum 8 characters"] =  util.ErrorMetaConfig{
		Tag:     "gte",
		Field:   "Password",
		Message: "password minimum must be 8 character",
	}
	if errResponse, errCount := util.GoValidator(&input, conf); errCount > 0 {
		util.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodPost, errResponse)
		return
	}
	data, status, err := h.service.ActiveUserSeekerRegister(&input)
	helpers.ErrUserRegisterHandler(ctx, status, err, data)
}
