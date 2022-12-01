package jobPostService

import (
	"github.com/KadirbekSharau/apprentice-backend/dto"
	"github.com/KadirbekSharau/apprentice-backend/models"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

/* Create Education Details Repository Service */
func (r *repository) CreateJobPost(input *dto.CreateEducationDetails) (*models.EducationDetails, string) {
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