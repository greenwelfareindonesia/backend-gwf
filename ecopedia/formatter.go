package ecopedia

import "time"

type EcopediaFormatter struct {
	ID int `json:"ID"`
	Judul     string `json:"Judul"`
	Subjudul  string `json:"SubJudul"`
	Deskripsi string `json:"Deskripsi"`
	Srcgambar string `json:"SrcGambar"`
	Referensi string `json:"Referensi"`
	FileName []string `json:"FileNames"`
	Comment []string `json:"Comment"`
	CreatedAt time.Time `json:"CreatedAt"`
    UpdatedAt time.Time `json:"UpdatedAt"`
	// Like []string `json:"like"`
}

func GetOneEcopediaFormat(ecopedia_submit Ecopedia) EcopediaFormatter {
	formatter := EcopediaFormatter{
		ID:  ecopedia_submit.ID,
		Judul:     ecopedia_submit.Judul,
		Subjudul:  ecopedia_submit.Subjudul,
		Deskripsi: ecopedia_submit.Deskripsi,
		Srcgambar: ecopedia_submit.Srcgambar,
		Referensi: ecopedia_submit.Referensi,
		// FileName: make([]string,len(ecopedia_submit.FileName)),
		// Comment: make([]string, len(ecopedia_submit.Comment)),
		CreatedAt: ecopedia_submit.CreatedAt,
	}

	for _, comment := range ecopedia_submit.Comment {
		formatter.Comment = append(formatter.Comment, comment.Comment)
	}

	for _, fileName := range ecopedia_submit.FileName {
		formatter.FileName = append(formatter.FileName, fileName.FileName)
	}

	return formatter
}

func FormatterGetAllEcopedia (ecopedia []Ecopedia) []EcopediaFormatter {
	newEcopediaGetFormatter := []EcopediaFormatter{}

	for _, newPounds := range ecopedia {
		newEcopediaFormatter := GetOneEcopediaFormat(newPounds)
		newEcopediaGetFormatter = append(newEcopediaGetFormatter, newEcopediaFormatter)
	}

	return newEcopediaGetFormatter
}
