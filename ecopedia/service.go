package ecopedia

type Service interface {
	FindAll() ([]Ecopedia, error)
	FindById(id int) (Ecopedia, error)
	CreateEcopedia(ecopedia EcopediaInput) (Ecopedia, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Ecopedia, error) {
	ecopedias, err := s.repository.FindAll()
	return ecopedias, err
	// return s.repository.FindAll(Ecopedia{})
}

func (s *service) FindById(id int) (Ecopedia, error) {
	return s.repository.FindById(id)
}

func (s *service) CreateEcopedia(ecopedia EcopediaInput) (Ecopedia, error) {
	newEcopedia := Ecopedia{}

	newEcopedia.Judul = ecopedia.Judul
	newEcopedia.Subjudul = ecopedia.Subjudul
	newEcopedia.Deskripsi = ecopedia.Deskripsi
	newEcopedia.Gambar = ecopedia.Gambar
	newEcopedia.Srcgambar = ecopedia.Srcgambar
	newEcopedia.Referensi = ecopedia.Referensi

	saveEcopedia, err := s.repository.Create(newEcopedia)

	if err != nil {
		return saveEcopedia, err
	}
	return saveEcopedia, nil
}
