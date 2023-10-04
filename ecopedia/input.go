package ecopedia

type EcopediaInput struct {
	Judul     string `form:"judul" binding:"required"`
	Subjudul  string `form:"subjudul"`
	Deskripsi string `form:"deskripsi" binding:"required"`
	Srcgambar string `form:"src_gambar"`
	Referensi string `form:"referensi"`
}

type EcopediaID struct {
	ID int `uri:"id" binding:"required"`
}	
