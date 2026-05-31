package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	Pool *pgxpool.Pool
}

func (r *UserRepository) CreateUser(username string, hashedPW string) (int, error) {
	var userID int

	err := r.Pool.QueryRow(
		context.Background(),
		`INSERT INTO users (username, hashed_pw)
		VALUES ($1, $2)
		RETURNING id
		`,
		username,
		hashedPW,
	).Scan(&userID)

	return userID, err
}
