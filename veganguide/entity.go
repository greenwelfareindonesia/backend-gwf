package veganguide

import "time"

type Veganguide struct {
	ID        int
	Judul     string
	Deskripsi string
	Body      string
	Gambar    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
