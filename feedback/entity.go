package feedback

import "time"

type Feedback struct {
	ID        int
	Email     string
	Text      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
