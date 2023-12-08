package veganguide

type Service interface {
	GetAllVeganguide(input int) ([]Veganguide, error)
	GetVeganguideByID(id int) (Veganguide, error)
	CreateVeganguide(veganguide VeganguideInput, FileName string) (Veganguide, error)
	DeleteVeganguide(ID int) (Veganguide, error)
	UpdateVeganguide(inputID VeganguideID, input VeganguideInput, FileName string) (Veganguide, error)
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
	if veganguides.ID == 0 {
		return veganguides, err
	}
	return veganguides, nil
}

func (s *service) CreateVeganguide(veganguide VeganguideInput, FileName string) (Veganguide, error) {
	newVeganguide := Veganguide{}

	newVeganguide.Judul = veganguide.Judul
	newVeganguide.Deskripsi = veganguide.Deskripsi
	newVeganguide.Body = veganguide.Body
	newVeganguide.Gambar = FileName

	saveVeganGuide, err := s.repository.Create(newVeganguide)

	if err != nil {
		return saveVeganGuide, err
	}
	return saveVeganGuide, nil
}

func (s *service) UpdateVeganguide(inputID VeganguideID, input VeganguideInput, FileName string) (Veganguide, error) {
	veganguide, err := s.repository.FindById(inputID.ID)
	if err != nil {
		return veganguide, err
	}

	veganguide.Judul = input.Judul
	veganguide.Deskripsi = input.Deskripsi
	veganguide.Body = input.Body
	veganguide.Gambar = FileName

	newVeganguide, err := s.repository.Update(veganguide)
	if err != nil {
		return newVeganguide, err
	}
	return newVeganguide, nil
}
