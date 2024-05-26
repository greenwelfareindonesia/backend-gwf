package entity

import "time"

type Veganguide struct {
	ID        int
	Slug      string
	Judul     string
	Deskripsi string
	Body      string
	Gambar    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
