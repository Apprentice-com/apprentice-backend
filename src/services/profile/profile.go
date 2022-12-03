package profileService

import (
	"github.com/KadirbekSharau/apprentice-backend/src/dto"
	model "github.com/KadirbekSharau/apprentice-backend/src/models"
)

type Service interface {
	GetSeekerProfile(input *dto.GetSeekerProfile) (*model.SeekerProfiles, int, string)
	CreateEducationDetails(input *dto.CreateEducationDetails) (*model.EducationDetails, int, string)
}

type service struct {
	repo *repository
}

func NewService(repo *repository) *service {
	return &service{repo: repo}
}

/* User Login Service */
func (s *service) GetSeekerProfile(input *dto.GetSeekerProfile) (*model.SeekerProfiles, int, string) {
	return s.repo.GetSeekerProfile(input)
}

/* Create EducationDetails Service */
func (s *service) CreateEducationDetails(input *dto.CreateEducationDetails) (*model.EducationDetails, int, string) {
	return s.repo.CreateEducationDetails(input)
}