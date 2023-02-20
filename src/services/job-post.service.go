package services

import (
	"github.com/KadirbekSharau/apprentice-backend/src/dto"
	"github.com/KadirbekSharau/apprentice-backend/src/models"
)

type JobPostService struct {
	repo *models.JobPostRepository
}

func NewJobPostService(repo *models.JobPostRepository) *JobPostService {
	return &JobPostService{repo: repo}
}

/* User Login Service */
func (s *JobPostService) CreateJobPost(input *dto.CreateJobPost) (*models.JobPost, int, string) {
	return s.repo.CreateJobPost(input)
}

/* Get All Job Posts */
func (s *JobPostService) GetAllJobPosts() (*[]models.JobPost, int, string) {
	return s.repo.GetAllJobPosts()
}

/* Get Job Post By ID */
func (s *JobPostService) GetJobPostByID(id string) (*models.JobPost, int, string) {
	return s.repo.GetJobPostByID(id)
}