package dto

type CreateArtikel struct {
	FullName       string `form:"fullName" binding:"required"`
	Email          string `form:"email" binding:"required"`
	Topic          string `form:"topic" binding:"required"`
	ArticleMessage string `form:"articleMessage" binding:"required"`
}

type GetArtikel struct {
	ID int `uri:"ID" binding:"required"`
}
