package contact

type ContactSubmissionInput struct {
	Name    string `json:"Name" `
	Email   string `json:"Email" binding:"required"`
	Subject string `json:"Subject" binding:"required"`
	Message string `json:"Message" binding:"required"`
}
