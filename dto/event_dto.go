package dto

type CreateEvents struct {
	Title        string `form:"title" binding:"required"`
	EventMessage string `form:"eventMessage" binding:"required"`
	Date         string `form:"date" `
	Location     string `form:"location" `
}

type GetEvent struct {
	ID int `uri:"id" binding:"required"`
}
