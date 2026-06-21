package db

import (
	"context"
	"secure-todo/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type TodoRepository struct {
	Pool *pgxpool.Pool
}

func (r *TodoRepository) CreateTodo(userID int, content string, dow int) (int, error) {
	var todoID int

	err := r.Pool.QueryRow(
		context.Background(),
		`INSERT INTO todos (user_id, content, dow)
 		VALUES ($1, $2, $3) RETURNING id`,
		userID,
		content,
		dow,
	).Scan(&todoID)

	return todoID, err
}

func (r *TodoRepository) GetTodosByUserID(userID int) ([]models.Todo, error) {
	var todos []models.Todo

	rows, err := r.Pool.Query(
		context.Background(),
		`SELECT id, content, dow, completed 
		FROM todos WHERE user_id=$1
		ORDER BY dow, id`,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var todo models.Todo

		err := rows.Scan(
			&todo.ID,
			&todo.Content,
			&todo.DOW,
			&todo.Completed,
		)
		if err != nil {
			return nil, err
		}

		todos = append(todos, todo)
	}

	return todos, rows.Err()
}

func (r *TodoRepository) UpdateTodo(userID int, todoID int, content string, dow int, completed bool) error {
	_, err := r.Pool.Exec(
		context.Background(),
		`UPDATE todos SET content=$1, dow=$2, completed=$3
		WHERE user_id=$4 and id=$5`,
		content,
		dow,
		completed,
		userID,
		todoID,
	)

	return err
}

func (r *TodoRepository) DeleteTodo(userID int, todoID int) error {
	_, err := r.Pool.Exec(
		context.Background(),
		`DELETE FROM todos WHERE user_id=$1 and id=$2`,
		userID,
		todoID,
	)

	return err
}
