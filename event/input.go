package event

type CreateEvents struct {
	Judul        string `form:"judul" binding:"required"`
	EventMessage string `form:"message" binding:"required"`
}

type GetEvent struct {
	ID int `uri:"id" binding:"required"`
}
