package dto

type InputKajian struct {
	Title       string `form:"title" binding:"required"`
	Description string `form:"description" binding:"required"`
}

type UpdateKajian struct {
	Title       string `form:"title"`
	Description string `form:"description"`
}
