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
func (r *repository) CreateSkillSet(input *dto.CreateSkillSet) (*models.SkillSet, int, string) {
	var entity models.SkillSet
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

func (r *repository) GetAllSkillSets() (*[]models.SkillSet, int, string) {
	var data []models.SkillSet
	if r.db.Model(&data).Debug().Select("*").Find(&data).Error != nil {
		return &[]models.SkillSet{}, http.StatusNotFound, "Data do not exist"
	}
	return &data, http.StatusOK, "nil"
}

func (r *repository) GetSkillSetByID(id string) (*models.SkillSet, int, string) {
	var data models.SkillSet
	if r.db.Model(&data).Debug().Select("*").Where("id = ?", id).Find(&data).RowsAffected != 1 {
		return nil, http.StatusNotFound, "Data not found"
	}
	return &data, http.StatusOK, "nil"
}