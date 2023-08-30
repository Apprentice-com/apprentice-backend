package repository

import (
	"context"
	"fmt"

	"github.com/KadirbekSharau/apprentice-backend/internal/models"
	"github.com/jmoiron/sqlx"
)

type jobPostRepository struct {
	db *sqlx.DB
}

func NewJobPostRepository(db *sqlx.DB) *jobPostRepository {
	return &jobPostRepository{
		db: db,
	}
}

func (r *jobPostRepository) CreateJobPost(ctx context.Context, jobPost *models.JobPost) error {
	query := `INSERT INTO job_posts 
			  (employer_id, location_id, title, description, level, experience_years, created_at, updated_at) VALUES 
			  (:employer_id, :location_id, :title, :description, :level, :experience_years, :created_at, :updated_at)`

	_, err := r.db.NamedExecContext(ctx, query, jobPost)
	if err != nil {
		fmt.Println("Error creating job post:", err)
		return err
	}
	return nil
}

func (r *jobPostRepository) GetAllJobPostsByEmployerID(ctx context.Context, employerID int) (*[]models.JobPost, error) {
	query := `SELECT * FROM job_posts WHERE employer_id = $1`
	var jobPosts []models.JobPost

	err := r.db.SelectContext(ctx, &jobPosts, query, employerID)
    if err != nil {
		fmt.Println("Error getting job posts:", err)
        return nil, err
    }
    
    return &jobPosts, nil
}

func (r *jobPostRepository) GetAllJobPosts(ctx context.Context) ([]models.JobPost, error) {
	query := `SELECT * FROM job_posts`
	var jobPosts []models.JobPost

	err := r.db.SelectContext(ctx, &jobPosts, query)
    if err != nil {
		fmt.Println("Error getting job posts:", err)
        return nil, err
    }
    
    return jobPosts, nil
}