package feedback

type FeedbackFormatter struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Text  string `json:"text"`
}

func FormatterFeedback(feedback_submission Feedback) FeedbackFormatter {
	formatter := FeedbackFormatter{
		ID:    feedback_submission.ID,
		Email: feedback_submission.Email,
		Text:  feedback_submission.Text,
	}
	return formatter
}
