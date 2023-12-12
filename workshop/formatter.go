package workshop

import "time"

type WorkshopFormatter struct {
	ID        int       `json:"id"`
	Slug      string    `json:"slug"`
	Title     string    `json:"title"`
	Image     string    `json:"image"`
	Desc      string    `json:"desc"`
	Date      string    `json:"date"`
	Url       string    `json:"url"`
	IsOpen    bool      `json:"is_open"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FormatterWorkshop(workshop Workshop) WorkshopFormatter {
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
		UpdatedAt: workshop.UpdatedAt,
	}
	return formatter
}
