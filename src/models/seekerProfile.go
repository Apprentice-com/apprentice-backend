package models

import (
	"net/http"

	"github.com/KadirbekSharau/apprentice-backend/src/dto"
	"gorm.io/gorm"
)

// SeekerProfile is database entity for each transaction
type SeekerProfile struct {
	UserID      uint 
	FirstName   string `gorm:"type:varchar(30)"`
	SecondName  string `gorm:"type:varchar(30)"`
	Description string `gorm:"type:varchar(1000)"`
	Resume      string `gorm:"type:varchar"`
}

type ProfileRepository struct {
	db *gorm.DB
}

func NewProfileRepository(db *gorm.DB) *ProfileRepository {
	return &ProfileRepository{db: db}
}

/* Get Profile Repository Service */
func (r *ProfileRepository) GetSeekerProfile(input *dto.GetSeekerProfile) (*SeekerProfile, int, string) {
	var profile SeekerProfile
	db := r.db.Model(&profile)

	if db.Debug().Select("*").Where("user_id = ?", input.UserID).Find(&profile).RowsAffected < 1 {
		return nil, http.StatusNotFound, "Data not found"
	}
	return &profile, http.StatusOK, "nil"
}
