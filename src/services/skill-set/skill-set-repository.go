package skillSetService

import (
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

/* Create Education Details Repository Service */
func (r *repository) CreateSkillSet(input *dto.CreateSkillSet) (*models.SkillSets, string) {
	var entity models.SkillSets
	db := r.db.Model(&entity)
	errorCode := make(chan string, 1)

	checkExist := db.Debug().Select("*").Where("name = ?", input.Name).Find(&entity)

	if checkExist.RowsAffected > 0 {
		errorCode <- "CREATE_CONFLICT_409"
		return &entity, <-errorCode
	}

	entity.Name = input.Name

	
	addNew := r.db.Debug().Create(&entity)

	db.Commit()

	if addNew.Error != nil {
		errorCode <- "CREATE_FAILED_403"
	} else {
		errorCode <- "nil"
	}

	return &entity, <-errorCode
}