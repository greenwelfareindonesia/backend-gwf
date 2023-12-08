package artikel

import "time"

type Artikel struct {
	ID            int
	Slug string
	FullName string
	Email string
	Topic string
	ArtikelMessage string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}