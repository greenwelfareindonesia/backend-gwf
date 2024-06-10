package entity

type Event struct {
	ID           int
	Slug         string
	Title        string
	EventMessage string
	FileName     string
	DefaultColumn
}
