package artikel

type CreateArtikel struct {
	FullName       string `form:"full_name" binding:"required"`
	Email          string `form:"email" binding:"required"`
	Topic          string `form:"topic" binding:"required"`
	ArtikelMessage string `form:"message" binding:"required"`
}

type GetArtikel struct {
	ID int `uri:"id" binding:"required"`
}
