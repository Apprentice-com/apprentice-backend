package employer

import (
	"context"

	"github.com/KadirbekSharau/apprentice-backend/internal/models"
)

type Repository interface {
	CreateEmployerProfile(ctx context.Context, profile *models.EmployerProfile) error
}
