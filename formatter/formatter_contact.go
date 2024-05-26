package formatter

import (
	"greenwelfare/entity"
	"time"
)

type ContactFormatter struct {
	ID int `json:"ID"`
	Name string `json:"name"`
	Slug string `json:"slug"`
	Email string `json:"email"`
	Subject string `json:"subject"`
	Message string `json:"message"`
	CreatedAt time.Time `json:"createdAt"`
}

func FormatterContact(contact_submission *entity.Contact) ContactFormatter {
	formatter := ContactFormatter{
		ID:      contact_submission.ID,
		Name:    contact_submission.Name,
		Slug:    contact_submission.Slug,
		Email:   contact_submission.Email,
		Subject: contact_submission.Subject,
		Message: contact_submission.Message,
	}
	return formatter
}
