package skillSetService

import (
	"github.com/KadirbekSharau/apprentice-backend/src/dto"
	model "github.com/KadirbekSharau/apprentice-backend/src/models"
)

type Service interface {
	CreateSkillSet(input *dto.CreateSkillSet) (*model.SkillSets, int, string)
}

type service struct {
	repo *repository
}

func NewService(repo *repository) *service {
	return &service{repo: repo}
}

/* Create Skill Set Service */
func (s *service) CreateSkillSet(input *dto.CreateSkillSet) (*model.SkillSets, int, string) {
	return s.repo.CreateSkillSet(input)
}