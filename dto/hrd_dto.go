package dto

type CreateHrdDTO struct {
	Nama        string `json:"nama"`
	Departement string `json:"departement"`
	Role        string `json:"role"`
	Status      string `json:"status"`
}

type UpdateHrdDTO struct {
	Nama        string `json:"nama"`
	Departement string `json:"departement"`
	Role        string `json:"role"`
	Status      string `json:"status"`
}
