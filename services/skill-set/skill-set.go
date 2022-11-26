package skillSetService

import (
	"github.com/KadirbekSharau/apprentice-backend/dto"
	model "github.com/KadirbekSharau/apprentice-backend/models"
)

type Service interface {
	CreateSkillSet(input *dto.CreateSkillSet) (*model.SkillSets, string)
}

type service struct {
	repo *repository
}

func NewService(repo *repository) *service {
	return &service{repo: repo}
}

/* Create Skill Set Service */
func (s *service) CreateSkillSet(input *dto.CreateSkillSet) (*model.SkillSets, string) {

	result, err := s.repo.CreateSkillSet(input)

	return result, err
}