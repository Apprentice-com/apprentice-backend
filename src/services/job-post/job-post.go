package jobPostService

import (
	"github.com/KadirbekSharau/apprentice-backend/src/dto"
	"github.com/KadirbekSharau/apprentice-backend/src/models"
)

type Service interface {
	GetAllJobPosts() (*[]models.JobPosts, string)
	CreateJobPost(input *dto.CreateJobPost) (*models.JobPosts, string)
}

type service struct {
	repo *repository
}

func NewService(repo *repository) *service {
	return &service{repo: repo}
}

/* User Login Service */
func (s *service) CreateJobPost(input *dto.CreateJobPost) (*models.JobPosts, string) {

	result, err := s.repo.CreateJobPost(input)

	return result, err
}

/* Get All Job Posts */
func (s *service) GetAllJobPosts() (*[]models.JobPosts, string) {
	result, err := s.repo.GetAllJobPosts()

	return result, err
}