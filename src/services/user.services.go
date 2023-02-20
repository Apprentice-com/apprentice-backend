package services

import (
	"github.com/KadirbekSharau/apprentice-backend/src/dto"
	model "github.com/KadirbekSharau/apprentice-backend/src/models"
)

type UserService struct {
	repo *model.UserRepository
}

func NewUserService(repo *model.UserRepository) *UserService {
	return &UserService{repo: repo}
}

/* User Login Service */
func (s *UserService) UserLogin(input *dto.InputLogin) (*model.Users, int, string) {
	return s.repo.UserLogin(input)
}

/* Active User Registration Service */
func (s *UserService) ActiveUserSeekerRegister(input *dto.InputUserSeekerRegister) (*model.Users, int, string) {
	input.IsActive = true
	input.UserType = 1
	return s.repo.UserRegister(input)
}

/* Admin User Registration Service */
func (s *UserService) AdminRegister(input *dto.InputUserSeekerRegister) (*model.Users, int, string) {
	input.IsActive = true
	input.UserType = 0
	return s.repo.UserRegister(input)
}