package authService

import (
	model "github.com/KadirbekSharau/apprentice-backend/models"
)

type Service interface {
	UserLogin(input *InputLogin) (*model.Users, string)
	ActiveUserRegister(input *InputUserRegister) (*model.Users, string)
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
func (s *service) ActiveUserRegister(input *InputUserRegister) (*model.Users, string) {

	users := model.Users{
		Email:    input.Email,
		Password: input.Password,
	}
	resultRegister, errRegister := s.repo.ActiveUserRegisterRepository(&users)

	return resultRegister, errRegister
}