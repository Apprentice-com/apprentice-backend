package auth

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/KadirbekSharau/apprentice-backend/internal/auth"
	"github.com/KadirbekSharau/apprentice-backend/internal/models"
	"github.com/dgrijalva/jwt-go/v4"
)

type TokenManager interface {
	ParseToken(ctx context.Context, accessToken string) (*models.User, error)
	SignTokenJWT(user *models.User) (string, error)
}

type AuthClaims struct {
	jwt.StandardClaims
	User *models.User `json:"user"`
}

type manager struct {
	signingKey     []byte
	expireDuration time.Duration
}

func NewTokenManager(signingKey []byte, expireDuration time.Duration) (TokenManager, error) {
	if signingKey == nil || expireDuration <= 0 {
		return nil, errors.New("empty signing key")
	}
	return &manager{
		signingKey:     signingKey,
		expireDuration: expireDuration,
	}, nil
}

// Parse auth token
func (m *manager) ParseToken(ctx context.Context, accessToken string) (*models.User, error) {
	token, err := jwt.ParseWithClaims(
		accessToken,
		&AuthClaims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return m.signingKey, nil
		})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*AuthClaims); ok && token.Valid {
		return claims.User, nil
	}

	return nil, auth.ErrInvalidAccessToken
}

// Sign Token with JWT
func (m *manager) SignTokenJWT(user *models.User) (string, error) {
	claims := AuthClaims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(m.expireDuration)),
		},
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(m.signingKey)
}
