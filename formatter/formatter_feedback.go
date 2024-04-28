package formatter

import (
	"greenwelfare/entity"
	"time"
)

type FeedbackFormatter struct {
	ID        int       `json:"ID"`
	Slug      string    `json:"slug"`
	Email     string    `json:"email"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"createdAt"`
}

func PostFormatterFeedback(feedback_submission *entity.Feedback) FeedbackFormatter {
	formatter := FeedbackFormatter{
		ID:        feedback_submission.ID,
		Slug:      feedback_submission.Slug,
		Email:     feedback_submission.Email,
		Text:      feedback_submission.Text,
		CreatedAt: feedback_submission.CreatedAt,
	}
	return formatter
}

