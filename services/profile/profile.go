package profileService

import (
	model "github.com/KadirbekSharau/apprentice-backend/models"
)

type Service interface {
	GetSeekerProfile(input *InputGetSeekerProfile) (*model.SeekerProfiles, string)
}

type service struct {
	repo *repository
}

func NewService(repo *repository) *service {
	return &service{repo: repo}
}

/* User Login Service */
func (s *service) GetSeekerProfile(input *InputGetSeekerProfile) (*model.SeekerProfiles, string) {

	result, err := s.repo.GetSeekerProfile(input)

	return result, err
}