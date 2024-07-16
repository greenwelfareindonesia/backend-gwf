package dto

type CreateEvents struct {
	Title        string `form:"title" binding:"required"`
	EventMessage string `form:"eventMessage" binding:"required"`
}

type GetEvent struct {
	ID int `uri:"id" binding:"required"`
}
