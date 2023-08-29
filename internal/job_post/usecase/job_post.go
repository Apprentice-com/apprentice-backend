package usecase

import (
	"context"
	"time"

	jobpost "github.com/KadirbekSharau/apprentice-backend/internal/job_post"
	"github.com/KadirbekSharau/apprentice-backend/internal/models"
)

type jobPostUseCase struct {
	repo jobpost.Repository
}

func NewJobPostUseCase(repo jobpost.Repository) *jobPostUseCase {
	return &jobPostUseCase{
		repo: repo,
	}
}

func (uc *jobPostUseCase) CreateJobPost(ctx context.Context, inp *jobpost.CreateJobPostInput) error {
	jobPost := &models.JobPost{
		EmployerID:      inp.EmployerID,
		LocationID:      inp.LocationID,
		Title:           inp.Title,
		Description:     inp.Description,
		Level:           inp.Level,
		ExperienceYears: inp.ExperienceYears,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
	return uc.repo.CreateJobPost(ctx, jobPost)
}
