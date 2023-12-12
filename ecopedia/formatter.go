package ecopedia

import "time"

type EcopediaFormatter struct {
	ID int `json:"ID"`
	Title     string `json:"Title"`
	Slug string `json:"Slug"`
	SubTitle  string `json:"SubTitle"`
	Description string `json:"Description"`
	SrcFile string `json:"SrcFile"`
	Reference string `json:"Reference"`
	FileName []string `json:"FileNames"`
	// Comment []string `json:"Comment"`
	CreatedAt time.Time `json:"CreatedAt"`
    UpdatedAt time.Time `json:"UpdatedAt"`
	// Like []string `json:"like"`
}

func GetOneEcopediaFormat(ecopedia_submit Ecopedia) EcopediaFormatter {
	formatter := EcopediaFormatter{
		ID:  ecopedia_submit.ID,
		Title:     ecopedia_submit.Title,
		Slug: ecopedia_submit.Slug,
		SubTitle:  ecopedia_submit.SubTitle,
		Description: ecopedia_submit.Description,
		SrcFile: ecopedia_submit.SrcFile,
		Reference: ecopedia_submit.Reference,
		// FileName: make([]string,len(ecopedia_submit.FileName)),
		// Comment: make([]string, len(ecopedia_submit.Comment)),
		CreatedAt: ecopedia_submit.CreatedAt,
	}

	// for _, comment := range ecopedia_submit.Comment {
	// 	formatter.Comment = append(formatter.Comment, comment.Comment)
	// }

	for _, fileName := range ecopedia_submit.FileName {
		formatter.FileName = append(formatter.FileName, fileName.FileName)
	}

	return formatter
}

// func FormatterGetAllEcopedia (ecopedia []Ecopedia) []EcopediaFormatter {
// 	newEcopediaGetFormatter := []EcopediaFormatter{}

// 	for _, newPounds := range ecopedia {
// 		newEcopediaFormatter := GetOneEcopediaFormat(newPounds)
// 		newEcopediaGetFormatter = append(newEcopediaGetFormatter, newEcopediaFormatter)
// 	}

// 	return newEcopediaGetFormatter
// }

// Fungsi untuk mengonversi slice Ecopedia menjadi slice EcopediaFormatter
func FormatterGetAllEcopedia(ecopedias []Ecopedia) []EcopediaFormatter {
    var ecopediaFormatters []EcopediaFormatter

    for _, ecopedia := range ecopedias {
        ecopediaFormatter := GetOneEcopediaFormat(ecopedia)
        ecopediaFormatters = append(ecopediaFormatters, ecopediaFormatter)
    }

    return ecopediaFormatters
}
