package product

import "time"

type Products struct {
	ID        int
	Name      string
	Price     int
	Stock     int
	CreatedAt time.Time
	UpdatedAt time.Time
}
