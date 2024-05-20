package dto

type EcopediaInput struct {
	Title       string `form:"Title" binding:"required"`
	SubTitle    string `form:"SubTitle"`
	Description string `form:"Description" binding:"required"`
	SrcFile     string `form:"SrcFile"`
	Reference   string `form:"Reference"`
}

type EcopediaID struct {
	ID int `uri:"ID" binding:"required"`
}
