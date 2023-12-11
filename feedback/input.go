package feedback

type FeedbackInput struct {
	Email string `json:"Email" binding:"required"`
	Text  string `json:"Text" binding:"required"`
}

type FeedbackID struct {
	ID int `uri:"id" binding:"required"`
}
