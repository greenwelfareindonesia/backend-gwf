package entity

import (
	"time"
)

type Contact struct {
	ID        int
	Slug      string
	Name      string
	Email     string
	Subject   string
	Message   string
	CreatedAt time.Time
}
