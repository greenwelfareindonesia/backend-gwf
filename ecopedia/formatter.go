package ecopedia

type EcopediaFormatter struct {
	ID int `json:"id"`
	Judul     string `json:"judul"`
	Subjudul  string `json:"subjudul"`
	Deskripsi string `json:"deskripsi"`
	Srcgambar string `json:"srcgambar"`
	Referensi string `json:"referensi"`
	Comment []string `json:"Comment"`
	// Like []string `json:"like"`

}

// type EcopediaComment struct {

// }

func FormatterEcopedia(ecopedia_submit Ecopedia) EcopediaFormatter {
	formatter := EcopediaFormatter{
		ID:  ecopedia_submit.ID,
		Judul:     ecopedia_submit.Judul,
		Subjudul:  ecopedia_submit.Subjudul,
		Deskripsi: ecopedia_submit.Deskripsi,
		Srcgambar: ecopedia_submit.Srcgambar,
		Referensi: ecopedia_submit.Referensi,
		Comment: make([]string, len(ecopedia_submit.Comment)),
	}

	for i, comment := range ecopedia_submit.Comment {
        formatter.Comment[i] = comment.Comment
    }

	return formatter
}
