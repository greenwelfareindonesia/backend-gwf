package artikel

import "time"

type Artikel struct {
	ID int
	Slug string
	FullName string
	Email string
	Topic string
	ArticleMessage string
	CreatedAt time.Time
	UpdatedAt time.Time
}
