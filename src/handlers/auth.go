package handlers

import (
	"net/http"

	"github.com/KadirbekSharau/apprentice-backend/src/dto"
	"github.com/KadirbekSharau/apprentice-backend/src/services/auth"
	"github.com/KadirbekSharau/apprentice-backend/src/util"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var roles = map[string]int{"seeker": 1, "employer": 2, "admin": 0} 
const expTime = 24*60*1

var AuthConfig = util.ErrorConfig{
	Options: map[string]util.ErrorMetaConfig{
		"Email required": {
			Tag:     "required",
			Field:   "Email",
			Message: "email is required on body",
		},
		"Email format not valid": {
			Tag:     "email",
			Field:   "Email",
			Message: "email format is not valid",
		},
		"Password required": {
			Tag:     "required",
			Field:   "Password",
			Message: "password is required on body",
		},
	},
}

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
	if errResponse, errCount := util.GoValidator(&input, AuthConfig.Options); errCount > 0 {
		util.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodPost, errResponse)
		return
	}
	data, status, err := h.service.UserLogin(&input)

	if status != 200 {
		util.APIResponse(ctx, err, status, http.MethodGet, data)
		return
	}
	accessTokenData := map[string]interface{}{"id": data.ID, "email": data.Email, "role": roles["user"]}
	h.createToken(accessTokenData, ctx, err)
}

/* Active User Register Handler */
func (h *authHandler) ActiveUserSeekerRegister(ctx *gin.Context) {
	var input dto.InputUserSeekerRegister
	ctx.ShouldBindJSON(&input)
	if errResponse, errCount := util.GoValidator(&input, AuthConfig.Options); errCount > 0 {
		util.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodPost, errResponse)
		return
	}
	data, status, err := h.service.ActiveUserSeekerRegister(&input)

	if status != 201 {
		util.APIResponse(ctx, err, status, http.MethodGet, data)
		return
	}
	accessTokenData := map[string]interface{}{"id": data.ID, "email": data.Email}
	h.createToken(accessTokenData, ctx, "Register new user account successfully")
}

func (h *authHandler) createToken(token map[string]interface{}, ctx *gin.Context, message string) {
	accessToken, errToken := util.Sign(token, "JWT_SECRET", expTime)
	if errToken != nil {
		defer logrus.Error(errToken.Error())
		util.APIResponse(ctx, "Generate accessToken failed", http.StatusBadRequest, http.MethodPost, nil)
		return
	}
	util.APIResponse(ctx, message, http.StatusCreated, http.MethodPost, map[string]string{"accessToken": accessToken})
} 