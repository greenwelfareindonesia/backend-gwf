package entity

import (
	"time"
)

type Workshop struct {
	ID        int
	Slug      string
	Title     string
	Image     string
	Desc      string
	Date      string
	Url       string
	IsOpen    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
