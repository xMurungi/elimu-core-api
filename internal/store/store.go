package store

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type User struct {
	ID          string
	FirebaseUID string
	Email       string
	DisplayName *string
	CreatedAt   time.Time
}

type Store struct {
	pool *pgxpool.Pool
}

func New(pool *pgxpool.Pool) *Store {
	return &Store{pool: pool}
}

func (s *Store) UpsertUser(ctx context.Context, firebaseUID, email string, displayName *string) (User, error) {
	// 1. put your SQL in a backtick string
	myquery := `INSERT INTO users (firebase_uid, email, display_name)
	VALUES ($1, $2, $3)
	ON CONFLICT (firebase_uid)
	DO UPDATE SET email = EXCLUDED.email,
				display_name = EXCLUDED.display_name
	RETURNING id, firebase_uid, email, display_name, created_at`

	// 2. var u User
	var u User

	// 3. call s.pool.QueryRow(ctx, <sql>, <$1>, <$2>, <$3>) then .Scan(...)
	err := s.pool.QueryRow(ctx, myquery, firebaseUID, email, displayName).Scan(&u.ID, &u.FirebaseUID, &u.Email, &u.DisplayName, &u.CreatedAt)
	//    (you already wrote the Scan arguments; update them for the new field names)

	// 4. if Scan returns an error, return User{} and wrap it:
	//    fmt.Errorf("upsert user: %w", err)
	// 5. otherwise return u, nil

	if err != nil {
		return User{}, fmt.Errorf("upsert user: %w", err)
	}

	return u, nil
}
