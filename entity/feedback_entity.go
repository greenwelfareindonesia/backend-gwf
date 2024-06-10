package entity

type Feedback struct {
	ID    int
	Slug  string
	Email string
	Text  string
	DefaultColumn
}
