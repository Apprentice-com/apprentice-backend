package usecase

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"time"

	"github.com/KadirbekSharau/apprentice-backend/internal/auth"
	"github.com/KadirbekSharau/apprentice-backend/internal/models"
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
	hashSalt       string
	signingKey     []byte
	expireDuration time.Duration
}

func NewAuthUseCase(repo auth.Repository, hashSalt string, signingKey []byte, expireDuration time.Duration) *authUseCase {
	return &authUseCase{
		repo:           repo,
		hashSalt:       hashSalt,
		signingKey:     signingKey,
		expireDuration: expireDuration,
	}
}

// Sign Up user
func (a *authUseCase) SignUp(ctx context.Context, inp *auth.SignUpInput) error {
	userExists, err := a.repo.UserExistsByEmail(ctx, inp.Email)
	if err != nil {
		return auth.ErrInvalidEmailFormat
	}

	if userExists {
		return auth.ErrEmailAlreadyExists
	}

	hashedPassword := a.hashPassword(inp.Password)

	user := &models.User{
		Email:    inp.Email,
		Password: hashedPassword,
		Role:     RoleAdmin,
		IsActive: true,
	}

	return a.repo.CreateUser(ctx, user)
}

// Sign In user
func (a *authUseCase) SignIn(ctx context.Context, inp *auth.SignInInput) (string, error) {
	user, err := a.repo.GetUserByEmail(ctx, inp.Email)
	if err != nil {
		return "", auth.ErrUserNotFound
	}

	hashedInputPassword := a.hashPassword(inp.Password)

	if hashedInputPassword != user.Password {
		return "", errors.New("incorrect password")
	}

	return a.signTokenJWT(user)
}

// Parse auth token
func (a *authUseCase) ParseToken(ctx context.Context, accessToken string) (*models.User, error) {
	token, err := jwt.ParseWithClaims(
		accessToken,
		&AuthClaims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return a.signingKey, nil
		})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*AuthClaims); ok && token.Valid {
		return claims.User, nil
	}

	return nil, auth.ErrInvalidAccessToken
}

// Hashing password with SHA 256
func (a *authUseCase) hashPassword(password string) string {
	pwd := sha256.New()
	pwd.Write([]byte(password))
	pwd.Write([]byte(a.hashSalt))
	return fmt.Sprintf("%x", pwd.Sum(nil))
}

// Sign Token with JWT
func (a *authUseCase) signTokenJWT(user *models.User) (string, error) {
	claims := AuthClaims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(a.expireDuration)),
		},
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(a.signingKey)
}
