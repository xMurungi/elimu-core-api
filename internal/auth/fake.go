package auth

import (
	"context"
)

type FakeVerifier struct {
	Tokens map[string]Identity
}

var _ TokenVerifier = FakeVerifier{}

func (f FakeVerifier) Verify(ctx context.Context, idToken string) (Identity, error) {
	// look the token up in f.Tokens
	// found -> return the identity and nil
	id, ok := f.Tokens[idToken]
	if !ok {
		return Identity{}, ErrInvalidToken
	}
	return id, nil

	// not found -> return Identity{} and ErrInvalidToken
	// return Identity{}, ErrInvalidToken
}
