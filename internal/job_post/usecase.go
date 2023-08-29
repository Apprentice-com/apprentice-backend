package jobpost

import (
	"context"
	//"github.com/KadirbekSharau/apprentice-backend/internal/models"
)

type UseCase interface {
	// Job Post related methods
	CreateJobPost(ctx context.Context, inp *CreateJobPostInput) error
	//GetAllJobPosts(ctx context.Context) ([]*models.JobPost, error)
	//GetJobPostDetails(ctx context.Context, jobPostID int) (*models.JobPost, error)
}
