package entity

type Contact struct {
	ID      int
	Slug    string
	Name    string
	Email   string
	Subject string
	Message string
	DefaultColumn
}
