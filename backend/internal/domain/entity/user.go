package entity

import "time"

type User struct {
	ID        int64
	UUID      string
	Name      string
	Email     string
	Password  string
	RoleID    int64
	CreatedAt time.Time
	UpdatedAt time.Time
}
