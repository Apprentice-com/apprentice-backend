package helpers

import (
	"net/http"

	"github.com/KadirbekSharau/apprentice-backend/src/models"
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

/* User Login Token Handler Function */
func UserLoginTokenHandler(ctx *gin.Context, status int, err string, data *models.Users) {
	if status != 200 {
		util.APIResponse(ctx, err, status, http.MethodGet, data)
		return
	}
	accessTokenData := map[string]interface{}{"id": data.ID, "email": data.Email, "role": roles["user"]}
	accessToken, errToken := util.Sign(accessTokenData, "JWT_SECRET", expTime)
	if errToken != nil {
		defer logrus.Error(errToken.Error())
		util.APIResponse(ctx, "Generate accessToken failed", http.StatusBadRequest, http.MethodPost, nil)
		return
	}
	util.APIResponse(ctx, err, status, http.MethodPost, map[string]string{"accessToken": accessToken})
}

/* User Registration Error Handler */
func ErrUserRegisterHandler(ctx *gin.Context, status int, err string, data *models.Users) {
	if status != 201 {
		util.APIResponse(ctx, err, status, http.MethodGet, data)
		return
	}
	accessTokenData := map[string]interface{}{"id": data.ID, "email": data.Email}
	accessToken, errToken := util.Sign(accessTokenData, "JWT_SECRET", 60)
	if errToken != nil {
		defer logrus.Error(errToken.Error())
		util.APIResponse(ctx, "Generate accessToken failed", http.StatusBadRequest, http.MethodPost, nil)
		return
	}
	util.APIResponse(ctx, "Register new user account successfully", http.StatusCreated, http.MethodPost, map[string]string{"accessToken": accessToken})
}