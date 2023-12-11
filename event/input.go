package event

type CreateEvents struct {
	Judul        string `form:"Judul" binding:"required"`
	EventMessage string `form:"Message" binding:"required"`
}

type GetEvent struct {
	ID int `uri:"id" binding:"required"`
}
