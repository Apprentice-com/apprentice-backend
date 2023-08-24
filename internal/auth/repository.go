package auth

import (
	"context"

	"github.com/KadirbekSharau/apprentice-backend/internal/models"
)

type Repository interface {
	CreateUser(ctx context.Context, user *models.User) error
	UserExistsByEmail(ctx context.Context, email string) (bool, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
}