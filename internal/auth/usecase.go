package auth

import (
	"context"

)

type UseCase interface {
	SignUpApplicant(ctx context.Context, inp *SignUpInput) error
	SignUpEmployer(ctx context.Context, inp *SignUpInput) error	
	SignIn(ctx context.Context, inp *SignInInput) (string, error)
}