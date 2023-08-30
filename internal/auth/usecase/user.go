package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/KadirbekSharau/apprentice-backend/internal/auth"
	"github.com/KadirbekSharau/apprentice-backend/internal/models"
	"github.com/KadirbekSharau/apprentice-backend/pkg/hash"
	token "github.com/KadirbekSharau/apprentice-backend/pkg/auth"
	"github.com/dgrijalva/jwt-go/v4"
)

const (
	RoleAdmin     = "Admin"
	RoleApplicant = "Applicant"
	RoleEmployer  = "Employer"
)

type AuthClaims struct {
	jwt.StandardClaims
	User *models.User `json:"user"`
}

type authUseCase struct {
	repo           auth.Repository
	hasher hash.PasswordHasher
	tokenManager token.TokenManager
}

func NewAuthUseCase(repo auth.Repository, hasher hash.PasswordHasher, tokenManager token.TokenManager) *authUseCase {
	return &authUseCase{
		repo:           repo,
		hasher: hasher,
		tokenManager: tokenManager,
	}
}

// Sign Up Applicant User
func (a *authUseCase) SignUpApplicant(ctx context.Context, inp *auth.SignUpInput) error {
	userExists, err := a.repo.UserExistsByEmail(ctx, inp.Email)
	if err != nil {
		return auth.ErrInvalidEmailFormat
	}

	if userExists {
		return auth.ErrEmailAlreadyExists
	}

	hashedPassword, err := a.hasher.Hash(inp.Password)
	if err != nil {
		return err
	}

	user := &models.User{
		Email:    inp.Email,
		Password: hashedPassword,
		Role:     RoleApplicant,
		IsActive: true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return a.repo.CreateUser(ctx, user)
}

// Sign Up Employer User
func (a *authUseCase) SignUpEmployer(ctx context.Context, inp *auth.SignUpInput) error {
	userExists, err := a.repo.UserExistsByEmail(ctx, inp.Email)
	if err != nil {
		return auth.ErrInvalidEmailFormat
	}

	if userExists {
		return auth.ErrEmailAlreadyExists
	}

	hashedPassword, err := a.hasher.Hash(inp.Password)
	if err != nil {
		return err
	}

	user := &models.User{
		Email:    inp.Email,
		Password: hashedPassword,
		Role:     RoleEmployer,
		IsActive: true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return a.repo.CreateUser(ctx, user)
}

// Sign In user
func (a *authUseCase) SignIn(ctx context.Context, inp *auth.SignInInput) (string, error) {
	user, err := a.repo.GetUserByEmail(ctx, inp.Email)
	if err != nil {
		return "", auth.ErrUserNotFound
	}

	hashedInputPassword, err := a.hasher.Hash(inp.Password)
	if err != nil {
		return "", err
	}

	if hashedInputPassword != user.Password {
		return "", errors.New("incorrect password")
	}

	return a.tokenManager.SignTokenJWT(user)
}