package feedback

import "time"

type FeedbackFormatter struct {
	ID        int       `json:"id"`
	Slug      string    `json:"slug"`
	Email     string    `json:"email"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"CreatedAt"`
}

func PostFormatterFeedback(feedback_submission Feedback) FeedbackFormatter {
	formatter := FeedbackFormatter{
		ID:        feedback_submission.ID,
		Slug:      feedback_submission.Slug,
		Email:     feedback_submission.Email,
		Text:      feedback_submission.Text,
		CreatedAt: feedback_submission.CreatedAt,
	}
	return formatter
}
