package skillSetService

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

/* Create Education Details Repository Service */
func (r *repository) CreateSkillSet(input *dto.CreateSkillSet) (*models.SkillSets, int, string) {
	var entity models.SkillSets
	db := r.db.Model(&entity)
	if db.Debug().Select("*").Where("name = ?", input.Name).Find(&entity).RowsAffected > 0 {
		return nil, http.StatusConflict, "Already exists"
	}

	entity.Name = input.Name
	if r.db.Debug().Create(&entity).Error != nil {
		return nil, http.StatusForbidden, "Create Failed"
	}	
	db.Commit()
	return &entity, http.StatusCreated, "nil"
}