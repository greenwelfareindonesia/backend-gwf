package contact

import "time"

type ContactFormatter struct {
	ID        int       `json:"Id"`
	Name      string    `json:"Name"`
	Slug      string    `json:"Slug"`
	Email     string    `json:"Email"`
	Subject   string    `json:"Subject"`
	Message   string    `json:"Message"`
	CreatedAt time.Time `json:"CreatedAt"`
}

func FormatterContact(contact_submission Contact) ContactFormatter {
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
