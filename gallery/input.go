package gallery

type InputGallery struct {
	Alt string `json:"alt" binding:"required"`
	//Likes string `json:"likes" binding:"required"`
}

type InputGalleryID struct {
	ID int `uri:"id" binding:"required"`
}
