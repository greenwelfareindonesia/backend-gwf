package dto

type VeganguideInput struct {
	Title     string `form:"Title" binding:"required"`
	Description string `form:"Description" binding:"required"`
	Body      string `form:"Body"`
}

type GetVeganguide struct {
	ID int `uri:"id" binding:"required"`
}