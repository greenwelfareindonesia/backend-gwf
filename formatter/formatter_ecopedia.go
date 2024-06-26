package formatter

import (
	"greenwelfare/entity"
	"time"
)

type EcopediaFormatter struct {
	ID int `json:"ID"`
	Title     string `json:"title"`
	Slug string `json:"slug"`
	SubTitle  string `json:"subTitle"`
	Description string `json:"description"`
	SrcFile string `json:"srcFile"`
	Reference string `json:"reference"`
	FileName []string `json:"fileNames"`
	// Comment []string `json:"comment"`
	CreatedAt time.Time `json:"createdAt"`
    UpdatedAt time.Time `json:"updatedAt"`
	// Like []string `json:"like"`
}

func GetOneEcopediaFormat(ecopedia_submit *entity.Ecopedia) EcopediaFormatter {
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
func FormatterGetAllEcopedia(ecopedias []*entity.Ecopedia) []EcopediaFormatter {
    var ecopediaFormatters []EcopediaFormatter

    for _, ecopedia := range ecopedias {
        ecopediaFormatter := GetOneEcopediaFormat(ecopedia)
        ecopediaFormatters = append(ecopediaFormatters, ecopediaFormatter)
    }

    return ecopediaFormatters
}
