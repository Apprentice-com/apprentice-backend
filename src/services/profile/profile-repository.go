package profileService

import (
	"net/http"

	"github.com/KadirbekSharau/apprentice-backend/src/dto"
	"github.com/KadirbekSharau/apprentice-backend/src/models"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

/* Get Profile Repository Service */
func (r *repository) GetSeekerProfile(input *dto.GetSeekerProfile) (*models.SeekerProfile, int, string) {
	var profile models.SeekerProfile
	db := r.db.Model(&profile)

	if db.Debug().Select("*").Where("user_id = ?", input.UserID).Find(&profile).RowsAffected < 1 {
		return nil, http.StatusNotFound, "Data not found"
	}
	return &profile, http.StatusOK, "nil"
}

/* Create Education Details Repository Service */
func (r *repository) CreateEducationDetails(input *dto.CreateEducationDetails) (*models.EducationDetail, int, string) {
	var ed models.EducationDetail
	db := r.db.Model(&ed)

	if db.Debug().Select("*").Where("institution_name = ? AND major = ? AND degree = ? AND user_id = ?", input.InstitutionName, input.Major, input.Degree, input.UserID).Find(&ed).RowsAffected > 0 {
		return nil, http.StatusConflict, "Already exists"
	}

	ed.InstitutionName = input.InstitutionName
	ed.Major = input.Major
	ed.Degree = input.Degree
	ed.StartDate = input.StartDate
	ed.EndDate = input.EndDate
	ed.UserID = input.UserID
	
	if r.db.Debug().Create(&ed).Error != nil {
		return nil, http.StatusForbidden, "Create failed"
	}
	db.Commit()
	return &ed, http.StatusCreated, "nil"
}