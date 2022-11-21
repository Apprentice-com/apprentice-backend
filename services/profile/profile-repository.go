package profileService

import (
	model "github.com/KadirbekSharau/apprentice-backend/models"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

/* Get Profile Repository Service */
func (r *repository) GetSeekerProfile(input *InputGetSeekerProfile) (*model.SeekerProfiles, string) {

	var profile model.SeekerProfiles
	db := r.db.Model(&profile)
	errorCode := make(chan string, 1)

	result := db.Debug().Select("*").Where("user_id = ?", input.UserID).Find(&profile)

	if result.RowsAffected < 1 {
		errorCode <- "CITY_NOT_FOUND_404"
		return &profile, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &profile, <-errorCode
}