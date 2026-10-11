package store

import (
	"context"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ptr(s string) *string { return &s }

func TestUpsertUser_Idempotent(t *testing.T) {
	url := os.Getenv("TEST_DATABASE_URL")
	if url == "" {
		t.Skip("TEST_DATABASE_URL not set")
	}

	ctx := context.Background()
	pool, err := pgxpool.New(ctx, url)
	if err != nil {
		t.Fatalf("connect: %v", err)
	}
	t.Cleanup(pool.Close) // registered first, so it runs LAST

	const uid = "test-uid-upsert-idempotent"
	const email = "upsert-test@example.com"

	t.Cleanup(func() {
		_, err := pool.Exec(context.Background(),
			`DELETE FROM users WHERE firebase_uid = $1`, uid)
		if err != nil {
			t.Errorf("cleanup: %v", err)
		}
	})

	s := New(pool)

	u1, err := s.UpsertUser(ctx, uid, email, ptr("First"))
	if err != nil {
		t.Fatalf("first upsert: %v", err)
	}

	// TODO: second upsert, same uid and email, ptr("Second"), into u2
	u2, err := s.UpsertUser(ctx, uid, email, ptr("Second"))
	if err != nil {
		t.Fatalf("Second upsert: %v", err)
	}
	// TODO: if the IDs differ, t.Errorf with both values
	if u1.ID != u2.ID {
		t.Errorf("ID's: first=%v, second=%v", u1.ID, u2.ID)
	}
	// TODO: if DisplayName is nil or not "Second", t.Errorf
	if u2.DisplayName == nil {
		t.Fatalf("second upsert: DisplayName is nil")
	}

	if *u2.DisplayName != "Second" {
		t.Errorf("DisplayName = %q, want %q", *u2.DisplayName, "second")
	}

}
