package dto

type InputGallery struct {
	Alt string `form:"Alt" binding:"required"`
}