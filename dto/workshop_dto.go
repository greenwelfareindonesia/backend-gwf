package dto

type CreateWorkshop struct {
	Title  string `form:"Title" binding:"required"`
	Description   string `form:"Description" binding:"required"`
	Date   string `form:"Date" binding:"required"`
	Url    string `form:"Url" binding:"required"`
	IsOpen bool   `form:"IsOpen" binding:"required"`
}


type GetWorkshop struct {
	ID int `uri:"id" binding:"required"`
}
