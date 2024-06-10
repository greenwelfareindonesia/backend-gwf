package entity

type Artikel struct {
	ID             int
	Slug           string
	FullName       string
	Email          string
	Topic          string
	ArticleMessage string
	DefaultColumn
}
