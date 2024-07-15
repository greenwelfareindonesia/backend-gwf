package dto

type InputGallery struct {
	Alt string `form:"alt" binding:"required"`
}