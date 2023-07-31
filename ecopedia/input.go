package ecopedia

type EcopediaInput struct {
	ID        string `json:"id"`
	Judul     string `json:"judul" binding:"required"`
	Subjudul  string `json:"subjudul"`
	Deskripsi string `json:"deskripsi" binding:"required"`
	Gambar    string `form:"filegambar" `
	Srcgambar string `json:"src_gambar"`
	Referensi string `json:"referensi"`
}
