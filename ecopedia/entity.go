package ecopedia

import "time"

type Ecopedia struct {
	ID        int 
	Judul     string
	Subjudul  string
	Deskripsi string
	Gambar    string
	Srcgambar string
	Referensi string
	CreatedAt time.Time
    UpdatedAt time.Time 
}