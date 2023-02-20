package services

import (
	"github.com/KadirbekSharau/apprentice-backend/src/dto"
	model "github.com/KadirbekSharau/apprentice-backend/src/models"
)

type ProfileService struct {
	repo *model.ProfileRepository
}

func NewProfileService(repo *model.ProfileRepository) *ProfileService {
	return &ProfileService{repo: repo}
}

/* Get Seeker Profile */
func (s *ProfileService) GetSeekerProfile(input *dto.GetSeekerProfile) (*model.SeekerProfile, int, string) {
	return s.repo.GetSeekerProfile(input)
}