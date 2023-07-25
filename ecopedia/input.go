package ecopedia

type EcopediaInput struct {
	Judul     string `form:"judul" binding:"required"`
	Subjudul  string `form:"subjudul"`
	Deskripsi string `form:"deskripsi" rbinding:"equired"`
	Gambar    string `form:"file_name" `
	Srcgambar string `form:"src_gambar"`
	Referensi string `json:"referensi" binding:"required"`
}