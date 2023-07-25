package ecopedia

import (
	"gorm.io/gorm"
)

type Ecopedia struct {
	gorm.Model
	Judul     string
	Subjudul  string
	Deskripsi string
	Gambar    string
	Srcgambar string
	Referensi string
}