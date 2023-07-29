package contact

type ContactFormatter struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Subject string `json:"subject"`
	Message string `json:"message"`
}

func FormatterContact(contact_submission Contact) ContactFormatter {
	formatter := ContactFormatter{
		ID:      contact_submission.ID,
		Name:    contact_submission.Name,
		Email:   contact_submission.Email,
		Subject: contact_submission.Subject,
		Message: contact_submission.Message,
	}
	return formatter
}
