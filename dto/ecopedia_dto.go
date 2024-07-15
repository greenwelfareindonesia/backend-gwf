package dto

type EcopediaInput struct {
	Title       string `form:"title" binding:"required"`
	SubTitle    string `form:"subTitle"`
	Description string `form:"description" binding:"required"`
	SrcFile     string `form:"srcFile"`
	Reference   string `form:"reference"`
}

type EcopediaID struct {
	ID int `uri:"ID" binding:"required"`
}
