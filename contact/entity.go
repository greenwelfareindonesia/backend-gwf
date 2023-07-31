package contact

import (
	"time"
)

type Contact struct {
	ID        int
	Name      string
	Email     string
	Subject   string
	Message   string
	CreatedAt time.Time
}
