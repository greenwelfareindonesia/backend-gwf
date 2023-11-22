package workshop

type CreateWorkshop struct {
	Title  string `json:"Title" binding:"required"`
	Desc   string `json:"Desc" binding:"required"`
	Date   string `json:"Date" binding:"required"`
	Url    string `json:"Url" binding:"required"`
	IsOpen bool   `json:"IsOpen" binding:"required"`
}

type UpdateWorkshop struct {
	Title  string `json:"Title"`
	Desc   string `json:"Desc"`
	Date   string `json:"Date"`
	Url    string `json:"Url"`
	IsOpen bool   `json:"IsOpen"`
}

type GetWorkshop struct {
	ID int `uri:"id" binding:"required"`
}
