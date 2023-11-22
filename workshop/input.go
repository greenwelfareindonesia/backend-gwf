package workshop

type CreateWorkshop struct {
	Title  string `json:"Title" binding:"required"`
	Desc   string `json:"Desc" binding:"required"`
	Date   string `json:"Date" binding:"required"`
	Url    string `json:"Url" binding:"required"`
	IsOpen bool   `json:"IsOpen" binding:"required"`
}


type GetWorkshop struct {
	ID int `uri:"id" binding:"required"`
}
