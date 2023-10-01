package workshop

type CreateWorkshop struct {
	Title  string `json:"title" binding:"required"`
	Image  string `json:"image" binding:"required"`
	Desc   string `json:"desc" binding:"required"`
	Date   string `json:"date" binding:"required"`
	Url    string `json:"url" binding:"required"`
	IsOpen bool   `json:"is_open" binding:"required"`
}

type UpdateWorkshop struct {
	Title  string `json:"title"`
	Image  string `json:"image"`
	Desc   string `json:"desc"`
	Date   string `json:"date"`
	Url    string `json:"url"`
	IsOpen bool   `json:"is_open"`
}

type GetWorkshop struct {
	ID int `uri:"id" binding:"required"`
}
