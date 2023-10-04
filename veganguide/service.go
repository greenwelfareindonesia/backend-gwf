package veganguide

type Service interface {
	GetAllVeganguide(input int) ([]Veganguide, error)
	GetVeganguideByID(id int) (Veganguide, error)
	CreateVeganguide(veganguide VeganguideInput, FileName string) (Veganguide, error)
	DeleteVeganguide(ID int) (Veganguide, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) DeleteVeganguide(ID int) (Veganguide, error) {
	veganguides, err := s.repository.FindById(ID)
	if err != nil {
		return veganguides, err
	}

	veganguide, err := s.repository.DeleteVeganguide(veganguides)
	if err != nil {
		return veganguide, err
	}
	return veganguide, nil
}

func (s *service) GetAllVeganguide(input int) ([]Veganguide, error) {
	veganguides, err := s.repository.FindAll()
	if err != nil {
		return veganguides, err
	}
	return veganguides, nil
}

func (s *service) GetVeganguideByID(id int) (Veganguide, error) {
	veganguides, err := s.repository.FindById(id)
	if err != nil {
		return veganguides, err
	}
	return veganguides, nil
}

func (s *service) CreateVeganguide(veganguide VeganguideInput, FileName string) (Veganguide, error) {
	newVeganguide := Veganguide{}

	newVeganguide.Judul = veganguide.Judul
	newVeganguide.Deskripsi = veganguide.Deskripsi
	newVeganguide.Body = veganguide.Body
	// newVeganguide.Gambar = veganguide.Gambar
	newVeganguide.Gambar = FileName

	saveVeganGuide, err := s.repository.Create(newVeganguide)

	if err != nil {
		return saveVeganGuide, err
	}
	return saveVeganGuide, nil
}
