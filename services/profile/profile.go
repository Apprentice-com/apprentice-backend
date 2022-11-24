package profileService

import (
	model "github.com/KadirbekSharau/apprentice-backend/models"
)

type Service interface {
	GetSeekerProfile(input *GetSeekerProfile) (*model.SeekerProfiles, string)
	CreateEducationDetails(input *CreateEducationDetails) (*model.EducationDetails, string)
}

type service struct {
	repo *repository
}

func NewService(repo *repository) *service {
	return &service{repo: repo}
}

/* User Login Service */
func (s *service) GetSeekerProfile(input *GetSeekerProfile) (*model.SeekerProfiles, string) {

	result, err := s.repo.GetSeekerProfile(input)

	return result, err
}

/* Create EducationDetails Service */
func (s *service) CreateEducationDetails(input *CreateEducationDetails) (*model.EducationDetails, string) {

	result, err := s.repo.CreateEducationDetails(input)

	return result, err
}