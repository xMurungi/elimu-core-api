package auth

import (
	"context"
	"errors"
	"testing"
)

func TestFakeVerifier_Verify(t *testing.T) {
	v := FakeVerifier{Tokens: map[string]Identity{
		"good-token": {
			FirebaseUID: "uid-1",
			Email:       "a@b.com",
		},
	}}

	tests := []struct {
		name    string
		token   string
		wantUID string
		wantErr error
	}{
		// TODO: row for "good-token": wantUID "uid-1", wantErr nil
		{
			name:    "good-token",
			token:   "good-token",
			wantUID: "uid-1",
		},
		// TODO: row for an unknown token: wantUID "", wantErr ErrInvalidToken
		{
			name:    "unknown-token",
			token:   "unknown-token",
			wantErr: ErrInvalidToken,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// call Verify, compare results
			got, err := v.Verify(context.Background(), tc.token)

			// TODO: if !errors.Is(err, tc.wantErr) { t.Errorf(...) }
			if !errors.Is(err, tc.wantErr) {
				t.Errorf("Error is: %v, want %v", err, tc.wantErr)
			}

			// TODO: if got.FirebaseUID != tc.wantUID { t.Errorf(...) }
			if got.FirebaseUID != tc.wantUID {
				t.Errorf("got %q, want %q", got.FirebaseUID, tc.wantUID)
			}

		})
	}
}
