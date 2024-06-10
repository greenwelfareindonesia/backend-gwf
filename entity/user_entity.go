package entity

type User struct {
	ID       int
	Slug     string
	Username string
	Email    string
	Password string
	Role     int
	DefaultColumn
}
