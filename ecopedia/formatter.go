package ecopedia

type EcopediaFormatter struct {
	Judul     string `form:"judul"`
	Subjudul  string `form:"subjudul"`
	Deskripsi string `form:"deskripsi"`
	Gambar    string `form:"gambar"`
	Srcgambar string `form:"srcgambar"`
	Referensi string `form:"referensi"`
}
