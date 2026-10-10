package store

import (
	"context"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ptr(s string) *string { return &s }

func TestUpsertUser_Indempotent(t *testing.T) {
	url := os.Getenv("TEST_DATABASE_URL")
	if url == "" {
		t.Skip("TEST_DATABASE_URL not set")
	}

	ctx := context.Background()
	pool, err := pgxpool.New(ctx, url)
	if err != nil {
		t.Fatalf("connect: %v", err)
	}
	defer pool.Close()

	s := New(pool)

	// TODO 1: register a cleanup that deletes the test user afterwards
	// TODO 2: call s.UpsertUser twice with the SAME firebase uid,
	//         using ptr("First") and then ptr("Second") as display names
	// TODO 3: assert the IDs are equal
	// TODO 4: assert the second call's DisplayName is "Second"
	
}
