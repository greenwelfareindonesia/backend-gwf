package feedback

import "time"

type FeedbackFormatter struct {
	ID    int    `json:"ID"`
	Email string `json:"Email"`
	Text  string `json:"Text"`
	CreatedAt time.Time `json:"CreatedAt"`
}

func PostFormatterFeedback(feedback_submission Feedback) FeedbackFormatter {
	formatter := FeedbackFormatter{
		ID:    feedback_submission.ID,
		Email: feedback_submission.Email,
		Text:  feedback_submission.Text,
		CreatedAt: feedback_submission.CreatedAt,
	}
	return formatter
}
