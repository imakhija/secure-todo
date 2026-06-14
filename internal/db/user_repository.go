package db

import (
	"context"
	"secure-todo/internal/models"

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

func (r *UserRepository) GetUserByUsername(username string) (*models.User, error) {
	user := &models.User{}

	err := r.Pool.QueryRow(
		context.Background(),
		`SELECT id, username, hashed_pw FROM users WHERE username = $1`,
		username,
	).Scan(&user.ID, &user.Username, &user.HashedPW)

	return user, err
}
