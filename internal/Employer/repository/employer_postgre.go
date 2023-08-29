package repository

import (
	"context"
	"fmt"

	"github.com/KadirbekSharau/apprentice-backend/internal/models"
	"github.com/jmoiron/sqlx"
)

type employerRepository struct {
	db *sqlx.DB
}

func NewEmployerRepository(db *sqlx.DB) *employerRepository {
	return &employerRepository{
		db: db,
	}
}

func (r *employerRepository) CreateEmployerProfile(ctx context.Context, profile *models.EmployerProfile) error {
	query := `INSERT INTO employer_profiles 
			  (user_id, company_id, first_name, last_name, created_at, updated_at) VALUES 
			  (:user_id, :company_id, :first_name, :last_name, :created_at, :updated_at)`

	_, err := r.db.NamedExecContext(ctx, query, profile)
	if err != nil {
		fmt.Println("Error creating employer profile:", err)
		return err
	}
	return nil
}
