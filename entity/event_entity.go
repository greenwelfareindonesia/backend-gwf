package entity

import "time"

type Event struct {
	ID           int
	Slug         string
	Title        string
	EventMessage string
	FileName     string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
