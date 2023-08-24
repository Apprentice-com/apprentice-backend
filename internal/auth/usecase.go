package auth

import (
	"context"

	"github.com/KadirbekSharau/apprentice-backend/internal/models"
)

type UseCase interface {
	SignUp(ctx context.Context, inp *SignUpInput) error
	SignIn(ctx context.Context, inp *SignInInput) (string, error)
	ParseToken(ctx context.Context, accessToken string) (*models.User, error)
}