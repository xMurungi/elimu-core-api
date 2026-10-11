package auth

import (
	"errors"
	fbauth "firebase.google.com/go/v4/auth"
	"testing"
)

func TestIdentityFromToken(t *testing.T) {
	tests := []struct {
		name    string
		tok     *fbauth.Token
		want    Identity
		wantErr error
	}{
		{
			name: "full token",
			tok: &fbauth.Token{
				UID:    "u1",
				Claims: map[string]interface{}{
					"email": "a@b.com", 
					"name": "Ann",
				},
			},
			want: Identity{FirebaseUID: "u1", Email: "a@b.com", DisplayName: "Ann"},
		},
		// rows: full token, no name, no email, empty UID, non-string email (123)
		{
			name: "full token 2",
			tok: &fbauth.Token{
				UID:    "",
				Claims: map[string]interface{}{
					"email": 123, 
					"name": "",
				},
			},
			want: Identity{FirebaseUID: "u1", Email: "a@b.com", DisplayName: "Ann"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := identityFromToken(tc.tok)

			// same two checks as before: errors.Is, then compare got and tc.want
			// Identity has only strings, so == on the whole struct works
			if !errors.Is(err, tc.wantErr) {
				t.Errorf("Error is: %v, want %v", err, tc.wantErr)
			}

			if got.DisplayName != tc.name {
				t.Errorf("got %q, want %q", got.FirebaseUID, tc.name)
			}
			if got.Email != tc.want.Email {
				t.Errorf("got %q, want %q", got.FirebaseUID, tc.want.Email)
			}
			if got.FirebaseUID != tc.want.FirebaseUID {
				t.Errorf("got %q, want %q", got.FirebaseUID, tc.want.FirebaseUID)
			}
		})
	}
}
