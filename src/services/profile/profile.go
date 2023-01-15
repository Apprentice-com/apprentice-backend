package profileService

import (
	"github.com/KadirbekSharau/apprentice-backend/src/dto"
	model "github.com/KadirbekSharau/apprentice-backend/src/models"
)

type Service interface {
	GetSeekerProfile(input *dto.GetSeekerProfile) (*model.SeekerProfile, int, string)
	CreateEducationDetails(input *dto.CreateEducationDetails) (*model.EducationDetail, int, string)
}

type service struct {
	repo *repository
}

func NewService(repo *repository) *service {
	return &service{repo: repo}
}

/* User Login Service */
func (s *service) GetSeekerProfile(input *dto.GetSeekerProfile) (*model.SeekerProfile, int, string) {
	return s.repo.GetSeekerProfile(input)
}

/* Create EducationDetails Service */
func (s *service) CreateEducationDetails(input *dto.CreateEducationDetails) (*model.EducationDetail, int, string) {
	return s.repo.CreateEducationDetails(input)
}