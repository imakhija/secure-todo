package models

import "time"

type User struct {
	ID        int
	Username  string
	HashedPW  string
	CreatedAt time.Time
}
