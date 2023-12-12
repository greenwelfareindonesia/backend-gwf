package event

type CreateEvents struct {
	Title        string `form:"Title" binding:"required"`
	EventMessage string `form:"	EventMessage" binding:"required"`
}

type GetEvent struct {
	ID int `uri:"id" binding:"required"`
}
