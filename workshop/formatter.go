package workshop

import "time"

type WorkshopFormatter struct {
	ID        int       `json:"ID"`
	Slug      string    `json:"Slug"`
	Title     string    `json:"Title"`
	Image     string    `json:"FileName"`
	Desc      string    `json:"Description"`
	Date      string    `json:"Date"`
	Url       string    `json:"Url"`
	IsOpen    bool      `json:"IsOpen"`
	CreatedAt time.Time `json:"CreatedAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
}

func PostFormatterWorkshop(workshop Workshop) WorkshopFormatter {
	formatter := WorkshopFormatter{
		ID:        workshop.ID,
		Slug:      workshop.Slug,
		Title:     workshop.Title,
		Image:     workshop.Image,
		Desc:      workshop.Desc,
		Date:      workshop.Date,
		Url:       workshop.Url,
		IsOpen:    workshop.IsOpen,
		CreatedAt: workshop.CreatedAt,
	}
	return formatter
}

func UpdateFormatterWorkshop(workshop Workshop) WorkshopFormatter {
	formatter := WorkshopFormatter{
		ID:        workshop.ID,
		Slug:      workshop.Slug,
		Title:     workshop.Title,
		Image:     workshop.Image,
		Desc:      workshop.Desc,
		Date:      workshop.Date,
		Url:       workshop.Url,
		IsOpen:    workshop.IsOpen,
		UpdatedAt: workshop.UpdatedAt,
	}
	return formatter
}