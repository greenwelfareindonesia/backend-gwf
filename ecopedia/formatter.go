package ecopedia

type EcopediaFormatter struct {
	Judul     string `form:"judul"`
	Subjudul  string `form:"subjudul"`
	Deskripsi string `form:"deskripsi"`
	Gambar    string `form:"gambar"`
	Srcgambar string `form:"srcgambar"`
	Referensi string `form:"referensi"`
}

func FormatterEcopedia(ecopedia_submit Ecopedia) EcopediaFormatter {
	formatter := EcopediaFormatter{
		Judul:     ecopedia_submit.Judul,
		Subjudul:  ecopedia_submit.Subjudul,
		Deskripsi: ecopedia_submit.Deskripsi,
	}
	return formatter
}
