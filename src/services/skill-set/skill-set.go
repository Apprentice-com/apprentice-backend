package skillSetService

import (
	"github.com/KadirbekSharau/apprentice-backend/src/dto"
	"github.com/KadirbekSharau/apprentice-backend/src/models"
	model "github.com/KadirbekSharau/apprentice-backend/src/models"
)

type Service interface {
	CreateSkillSet(input *dto.CreateSkillSet) (*model.SkillSet, int, string)
	GetAllSkillSets() (*[]models.SkillSet, int, string)
	GetSkillSetByID(id string) (*models.SkillSet, int, string)
}

type service struct {
	repo *repository
}

func NewService(repo *repository) *service {
	return &service{repo: repo}
}

/* Create Skill Set */
func (s *service) CreateSkillSet(input *dto.CreateSkillSet) (*model.SkillSet, int, string) {
	return s.repo.CreateSkillSet(input)
}

/* Get All Skill Sets */
func (s *service) GetAllSkillSets() (*[]models.SkillSet, int, string) {
	return s.repo.GetAllSkillSets()
}

/* Get Skill Set By ID */
func (s *service) GetSkillSetByID(id string) (*models.SkillSet, int, string) {
	return s.repo.GetSkillSetByID(id)
}