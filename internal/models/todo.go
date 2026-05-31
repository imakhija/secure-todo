package models

import "time"

type Todo struct {
	ID        int
	UserID    int
	Content   string
	DOW       int
	Completed bool
	CreatedAt time.Time
}
