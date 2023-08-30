package jobpost

import (
	"context"

	"github.com/KadirbekSharau/apprentice-backend/internal/models"
)

type Repository interface {
	CreateJobPost(ctx context.Context, jobPost *models.JobPost) error
	GetAllJobPostsByEmployerID(ctx context.Context, employerID int) (*[]models.JobPost, error)
	GetAllJobPosts(ctx context.Context) ([]models.JobPost, error)
	//GetJobPostDetails(ctx context.Context, jobPostID int) (*models.JobPost, error)
}
