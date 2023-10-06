package feedback

type FeedbackInput struct {
	Email string `json:"email" binding:"required"`
	Text  string `json:"text" binding:"required"`
}

type FeedbackID struct {
	ID int `uri:"id" binding:"required"`
}
