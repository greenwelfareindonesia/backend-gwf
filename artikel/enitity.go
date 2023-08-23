package artikel

import "time"

type Artikel struct {
	ID            int
	FullName string
	Email string
	Topic string
	ArtikelMessage string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}