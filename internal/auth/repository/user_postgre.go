package repository

import (
	"context"
	"fmt"

	"github.com/KadirbekSharau/apprentice-backend/internal/auth"
	"github.com/KadirbekSharau/apprentice-backend/internal/models"
	"github.com/jmoiron/sqlx"
)

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) CreateUser(ctx context.Context, user *models.User) error {
	query := `INSERT INTO users (email, password, role, is_active, created_at, updated_at) 
	          VALUES (:email, :password, :role, :is_active, :created_at, :updated_at)`
	_, err := r.db.NamedExecContext(ctx, query, user)
	if err != nil {
		fmt.Println("Error creating user:", err)
		return err
	}
	return nil
}

func (r *userRepository) UserExistsByEmail(ctx context.Context, email string) (bool, error) {
	query := `SELECT COUNT(*) FROM users WHERE email = $1`
	var count int
	err := r.db.GetContext(ctx, &count, query, email)

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	query := "SELECT * FROM users WHERE email = $1"
	user := &models.User{}
	err := r.db.GetContext(ctx, user, query, email)

	if err != nil {
		fmt.Println("User not found or invalid credentials")
		return nil, auth.ErrUserNotFound
	}

	return user, nil
}
