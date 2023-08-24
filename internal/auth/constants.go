package auth

import "errors"

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidAccessToken = errors.New("invalid access token")
	ErrEmailAlreadyExists = errors.New("email already exists")
    ErrInvalidEmailFormat = errors.New("invalid email format")
    ErrInvalidPassword    = errors.New("invalid password")
)