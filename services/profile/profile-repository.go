package profileService

import (
	"github.com/KadirbekSharau/apprentice-backend/models"
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
func (r *repository) GetSeekerProfile(input *GetSeekerProfile) (*model.SeekerProfiles, string) {

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

/* Create Education Details Repository Service */
func (r *repository) CreateEducationDetails(input *CreateEducationDetails) (*model.EducationDetails, string) {
	var ed models.EducationDetails
	db := r.db.Model(&ed)
	errorCode := make(chan string, 1)

	checkExist := db.Debug().Select("*").Where("institution_name = ? AND major = ? AND degree = ? AND user_id = ?", input.InstitutionName, input.Major, input.Degree, input.UserID).Find(&ed)

	if checkExist.RowsAffected > 0 {
		errorCode <- "CREATE_CONFLICT_409"
		return &ed, <-errorCode
	}

	ed.InstitutionName = input.InstitutionName
	ed.Major = input.Major
	ed.Degree = input.Degree
	ed.StartDate = input.StartDate
	ed.EndDate = input.EndDate
	ed.UserID = input.UserID
	
	addNew := r.db.Debug().Create(&ed)

	db.Commit()

	if addNew.Error != nil {
		errorCode <- "CREATE_FAILED_403"
	} else {
		errorCode <- "nil"
	}

	return &ed, <-errorCode
}