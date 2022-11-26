package helpers

import (
	"net/http"

	"github.com/KadirbekSharau/apprentice-backend/models"
	"github.com/KadirbekSharau/apprentice-backend/util"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var roles = map[string]int{"seeker": 1, "employer": 2, "admin": 0} 
const secret_key string = "JWT_SECRET"
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
func UserLoginTokenHandler(ctx *gin.Context, errLogin string, resultLogin *models.Users) {
	switch errLogin {

	case "LOGIN_NOT_FOUND_404":
		util.APIResponse(ctx, "User account is not registered", http.StatusNotFound, http.MethodPost, nil)
		return

	case "LOGIN_NOT_ACTIVE_403":
		util.APIResponse(ctx, "User account is not active", http.StatusForbidden, http.MethodPost, nil)
		return

	case "LOGIN_WRONG_PASSWORD_403":
		util.APIResponse(ctx, "Username or password is wrong", http.StatusForbidden, http.MethodPost, nil)
		return

	default:
		accessTokenData := map[string]interface{}{"id": resultLogin.ID, "email": resultLogin.Email, "role": roles["user"]}
		accessToken, errToken := util.Sign(accessTokenData, secret_key, expTime)

		if errToken != nil {
			defer logrus.Error(errToken.Error())
			util.APIResponse(ctx, "Generate accessToken failed", http.StatusBadRequest, http.MethodPost, nil)
			return
		}

		util.APIResponse(ctx, "Login successfully", http.StatusOK, http.MethodPost, map[string]string{"accessToken": accessToken})
	}
}

/* User Registration Error Handler */
func ErrUserRegisterHandler(resultRegister *models.Users, ctx *gin.Context, errRegister string) {
	switch errRegister {

	case "REGISTER_CONFLICT_409":
		util.APIResponse(ctx, "Email already exist", http.StatusConflict, http.MethodPost, nil)
		return

	case "REGISTER_FAILED_403":
		util.APIResponse(ctx, "Register new User account failed", http.StatusForbidden, http.MethodPost, nil)
		return

	default:
		accessTokenData := map[string]interface{}{"id": resultRegister.ID, "email": resultRegister.Email}
		accessToken, errToken := util.Sign(accessTokenData, util.GodotEnv("JWT_SECRET"), 60)

		if errToken != nil {
			defer logrus.Error(errToken.Error())
			util.APIResponse(ctx, "Generate accessToken failed", http.StatusBadRequest, http.MethodPost, nil)
			return
		}

		util.APIResponse(ctx, "Register new user account successfully", http.StatusCreated, http.MethodPost, map[string]string{"accessToken": accessToken})
	}
}