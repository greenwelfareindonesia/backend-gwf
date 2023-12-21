package artikel

type CreateArtikel struct {
	FullName       string `form:"FullName" binding:"required"`
	Email          string `form:"Email" binding:"required"`
	Topic          string `form:"Topic" binding:"required"`
	ArticleMessage string `form:"ArticleMessage" binding:"required"`
}

type GetArtikel struct {
	ID int `uri:"ID" binding:"required"`
}
