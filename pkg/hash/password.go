package hash

import (
	"crypto/sha256"
	"fmt"
)

type PasswordHasher interface {
	Hash(password string) (string, error)
}

type hasherSHA256 struct {
	salt string
}

func NewHasherSHA256(salt string) PasswordHasher {
	return &hasherSHA256{
		salt: salt,
	}
}

// Hashing password with SHA 256
func (h *hasherSHA256) Hash(password string) (string, error) {
	pwd := sha256.New()
	if _, err := pwd.Write([]byte(password)); err != nil {
		return "", err
	}
	pwd.Write([]byte(h.salt))
	return fmt.Sprintf("%x", pwd.Sum(nil)), nil
}