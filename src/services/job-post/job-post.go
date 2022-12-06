package jobPostService

import (
	"github.com/KadirbekSharau/apprentice-backend/src/dto"
	"github.com/KadirbekSharau/apprentice-backend/src/models"
)

type Service interface {
	GetAllJobPosts() (*[]models.JobPosts, int, string)
	CreateJobPost(input *dto.CreateJobPost) (*models.JobPosts, int, string)
	GetJobPostByID(id string) (*models.JobPosts, int, string)
}

type service struct {
	repo *repository
}

func NewService(repo *repository) *service {
	return &service{repo: repo}
}

/* User Login Service */
func (s *service) CreateJobPost(input *dto.CreateJobPost) (*models.JobPosts, int, string) {
	return s.repo.CreateJobPost(input)
}

/* Get All Job Posts */
func (s *service) GetAllJobPosts() (*[]models.JobPosts, int, string) {
	return s.repo.GetAllJobPosts()
}

/* Get Job Post By ID */
func (s *service) GetJobPostByID(id string) (*models.JobPosts, int, string) {
	return s.repo.GetJobPostByID(id)
}