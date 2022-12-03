package authService

import (
	"github.com/KadirbekSharau/apprentice-backend/src/dto"
	model "github.com/KadirbekSharau/apprentice-backend/src/models"
)

type Service interface {
	UserLogin(input *dto.InputLogin) (*model.Users, int, string)
	ActiveUserSeekerRegister(input *dto.InputUserSeekerRegister) (*model.Users, int, string)
	AdminRegister(input *dto.InputUserSeekerRegister) (*model.Users, int, string)
}

type service struct {
	repo *repository
}

func NewService(repo *repository) *service {
	return &service{repo: repo}
}

/* User Login Service */
func (s *service) UserLogin(input *dto.InputLogin) (*model.Users, int, string) {
	return s.repo.UserLogin(input)
}

/* Active User Registration Service */
func (s *service) ActiveUserSeekerRegister(input *dto.InputUserSeekerRegister) (*model.Users, int, string) {
	return s.repo.ActiveUserSeekerRegister(input)
}

/* Admin User Registration Service */
func (s *service) AdminRegister(input *dto.InputUserSeekerRegister) (*model.Users, int, string) {
	return s.repo.AdminRegister(input)
}