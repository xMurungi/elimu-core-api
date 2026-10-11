package auth

import (
	"context"
	"errors"
)

type Identity struct {
	FirebaseUID string
	Email       string
	DisplayName string
}

type TokenVerifier interface {
	Verify(ctx context.Context, idToken string) (Identity, error)
}

var ErrInvalidToken = errors.New("invalid token")
