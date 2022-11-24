package helpers

import "github.com/KadirbekSharau/apprentice-backend/util"

var Configs = util.ErrorConfig{
	Options: map[string]util.ErrorMetaConfig{
		"ID required": {
			Tag:     "required",
			Field:   "ID",
			Message: "id is required on param",
		},
	},
}