package models

import "time"

type Todo struct {
	ID        int       `json:"-"`
	UserID    int       `json:"-"`
	Content   string    `json:"content"`
	DOW       int       `json:"dow"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"-"`
}
