package user

import (
	"time"
)

type User struct {
	ID        int
	Slug      string
	Username  string
	Email     string
	Password  string
	Role      int
	CreatedAt time.Time
	UpdatedAt time.Time
}
