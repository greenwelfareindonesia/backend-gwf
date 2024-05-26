package entity

import "time"

type Feedback struct {
	ID        int
	Slug      string
	Email     string
	Text      string
	CreatedAt time.Time
	UpdatedAt time.Time
}