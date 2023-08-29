package usecase

import (
	"context"
	"time"

	"github.com/KadirbekSharau/apprentice-backend/internal/employer"
	"github.com/KadirbekSharau/apprentice-backend/internal/models"
)


type employerUseCase struct {
	repo employer.Repository
}

func NewEmployerUseCase(repo employer.Repository) *employerUseCase {
	return &employerUseCase{
		repo: repo,
	}
}

func (uc *employerUseCase) CreateEmployerProfile(ctx context.Context, inp *employer.CreateEmployerProfileInput) error {
	profile := &models.EmployerProfile{
		UserID: inp.UserID,
		CompanyID: inp.CompanyID,
		FirstName: inp.FirstName,
		LastName: inp.LastName,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return uc.repo.CreateEmployerProfile(ctx, profile)
}