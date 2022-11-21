package authService

import (
	model "github.com/KadirbekSharau/apprentice-backend/models"
)

type Service interface {
	UserLogin(input *InputLogin) (*model.Users, string)
	ActiveUserSeekerRegister(input *InputUserSeekerRegister) (*model.Users, string)
}

type service struct {
	repo *repository
}

func NewService(repo *repository) *service {
	return &service{repo: repo}
}

/* User Login Service */
func (s *service) UserLogin(input *InputLogin) (*model.Users, string) {

	user := model.Users{
		Email:    input.Email,
		Password: input.Password,
	}

	resultLogin, errLogin := s.repo.UserLogin(&user)

	return resultLogin, errLogin
}

/* Active User Registration Service */
func (s *service) ActiveUserSeekerRegister(input *InputUserSeekerRegister) (*model.Users, string) {

	resultRegister, errRegister := s.repo.ActiveUserSeekerRegisterRepository(input)

	return resultRegister, errRegister
}