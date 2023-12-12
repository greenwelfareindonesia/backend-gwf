package feedback

import "time"

type FeedbackFormatter struct {
<<<<<<< Updated upstream
=======
<<<<<<< HEAD
	ID    int    `json:"id"`
	Slug  string `json:"slug"`
	Email string `json:"email"`
	Text  string `json:"text"`
=======
>>>>>>> Stashed changes
	ID    int    `json:"ID"`
	Email string `json:"Email"`
	Text  string `json:"Text"`
	CreatedAt time.Time `json:"CreatedAt"`
<<<<<<< Updated upstream
=======
>>>>>>> 0075203b2f39f41648b074b4c80756123471ee58
>>>>>>> Stashed changes
}

func PostFormatterFeedback(feedback_submission Feedback) FeedbackFormatter {
	formatter := FeedbackFormatter{
		ID:    feedback_submission.ID,
		Slug:  feedback_submission.Slug,
		Email: feedback_submission.Email,
		Text:  feedback_submission.Text,
		CreatedAt: feedback_submission.CreatedAt,
	}
	return formatter
}
