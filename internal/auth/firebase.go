package auth

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go/v4"
	fbauth "firebase.google.com/go/v4/auth"
)

func identityFromToken(tok *fbauth.Token) (Identity, error) {
	if tok.UID == "" {
		return Identity{}, fmt.Errorf("missing uid: %w", ErrInvalidToken)
	}

	email, ok := tok.Claims["email"].(string)
	if !ok || email == ""{
		return Identity{}, fmt.Errorf("missing email claim: %w", ErrInvalidToken)
	}

	name, _ := tok.Claims["name"].(string)

}

type FirebaseVerifier struct {
	client *fbauth.Client
}

var _ TokenVerifier = (*FirebaseVerifier)(nil)

func NewFirebaseVerifier(ctx context.Context, projectID string) (*FirebaseVerifier, error) {
	// 1. app, err := firebase.NewApp(ctx, &firebase.Config{ProjectID: projectID})
	app, err := firebase.NewApp(ctx, &firebase.Config{ProjectID: projectID})
	if err != nil {
		return nil, fmt.Errorf("init firebase app: %w", err)
	}

	// 2. client, err := app.Auth(ctx)
	client, err := app.Auth(ctx)
	if err != nil {
		return nil, fmt.Errorf("init firebase client: %w", err)
	}

	// 3. wrap errors with fmt.Errorf("...: %w", err), return &FirebaseVerifier{client: client}, nil
	fmt.Errorf("...: %w", err), return &FirebaseVerifier{client: client}, nil
	
}

func (f *FirebaseVerifier) Verify(ctx context.Context, idToken string) (Identity, error) {
	// 1. tok, err := f.client.VerifyIDToken(ctx, idToken)
	tok, err := f.client.VerifyIDToken(ctx, idToken)

	// 2. if err != nil: decide what kind of failure it is (see below)
	if err != nil {
		if IsIDTokenExpired {
			return Identity{}, fmt.Errorf("verify id token: %w", ErrInvalidToken)
		}
		if IsIDTokenInvalid {
			return Identity{}, fmt.Errorf("verify id token: %w", ErrInvalidToken)
		}
		if IsIDTokenRevoked {
			return Identity{}, fmt.Errorf("verify id token: %w", ErrInvalidToken)
		}
		return Identity{}, fmt.Errorf("verify id token: %w", err)
	}

	// 3. return identityFromToken(tok)
	return identityFromToken(tok)
}
