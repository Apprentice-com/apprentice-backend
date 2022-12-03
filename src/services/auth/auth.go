package authService

import (
	"github.com/KadirbekSharau/apprentice-backend/src/dto"
	model "github.com/KadirbekSharau/apprentice-backend/src/models"
)

type Service interface {
	UserLogin(input *dto.InputLogin) (*model.Users, string)
	ActiveUserSeekerRegister(input *dto.InputUserSeekerRegister) (*model.Users, string)
	AdminRegister(input *dto.InputUserSeekerRegister) (*model.Users, string)
}

type service struct {
	repo *repository
}

func NewService(repo *repository) *service {
	return &service{repo: repo}
}

/* User Login Service */
func (s *service) UserLogin(input *dto.InputLogin) (*model.Users, string) {

	user := model.Users{
		Email:    input.Email,
		Password: input.Password,
	}

	resultLogin, errLogin := s.repo.UserLogin(&user)

	return resultLogin, errLogin
}

/* Active User Registration Service */
func (s *service) ActiveUserSeekerRegister(input *dto.InputUserSeekerRegister) (*model.Users, string) {

	resultRegister, errRegister := s.repo.ActiveUserSeekerRegister(input)

	return resultRegister, errRegister
}

/* Admin User Registration Service */
func (s *service) AdminRegister(input *dto.InputUserSeekerRegister) (*model.Users, string) {

	resultRegister, errRegister := s.repo.AdminRegister(input)

	return resultRegister, errRegister
}