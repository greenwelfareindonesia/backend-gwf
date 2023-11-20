package veganguide

type VeganguideInput struct {
	Judul     string `form:"judul" binding:"required"`
	Deskripsi string `form:"deskripsi" binding:"required"`
	Body      string `form:"body"`
}

type VeganguideID struct {
	ID int `uri:"id" binding:"required"`
}