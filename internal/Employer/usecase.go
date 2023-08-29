package employer

import (
	"context"
	//"github.com/KadirbekSharau/apprentice-backend/internal/models"
)

type UseCase interface {
	// Employer-related methods
	CreateEmployerProfile(ctx context.Context, inp *CreateEmployerProfileInput) error
	//CreateJobPost(ctx context.Context, inp *CreateJobPostInput) error
	//GetAllJobPosts(ctx context.Context) ([]*models.JobPost, error)
	//GetJobPostDetails(ctx context.Context, jobPostID int) (*models.JobPost, error)
	//GetApplicantsForJobPost(ctx context.Context, jobPostID int) ([]*models.ApplicantProfile, error)
	//UpdateApplicantStatus(ctx context.Context, jobPostID int, applicantID int, status string) error
}
