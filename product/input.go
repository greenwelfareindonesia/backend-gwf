package product

type ProductInput struct {
	Title string `json:"title" binding:"required"`
	Price int    `json:"price" binding:"required"`
	Stock int    `json:"stock" binding:"required"`
}

type GetinputProductID struct {
	ID int `uri:"id" binding:"required"`
}

type GetCategoryID struct {
	ID int `uri:"categoryID" binding:"required"`
}
