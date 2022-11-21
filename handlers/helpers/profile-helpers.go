package helpers

import "github.com/KadirbekSharau/apprentice-backend/util"

var Configs = util.ErrorConfig{
	Options: []util.ErrorMetaConfig{
		{
			Tag:     "required",
			Field:   "ID",
			Message: "id is required on param",
		},
		{
			Tag:     "email",
			Field:   "Email",
			Message: "email format is not valid",
		},
		{
			Tag:     "required",
			Field:   "Password",
			Message: "password is required on body",
		},
	},
}