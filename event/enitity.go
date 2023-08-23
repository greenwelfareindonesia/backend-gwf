package event

import "time"

type Event struct {
	ID            int
	Judul string
	EventMessage string
	FileName string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}