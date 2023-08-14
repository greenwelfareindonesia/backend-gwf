package ecopedia

type Service interface {
	GetAllEcopedia(input int) ([]Ecopedia, error)
	FindById(id int) (Ecopedia, error)
	CreateEcopedia(ecopedia EcopediaInput, FileName string) (Ecopedia, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAllEcopedia(input int) ([]Ecopedia, error) {
	ecopedias, err := s.repository.FindAll()
	if err != nil {
		return ecopedias, err
	}
	return ecopedias, nil
}

func (s *service) FindById(id int) (Ecopedia, error) {
	return s.repository.FindById(id)
}

func (s *service) CreateEcopedia(ecopedia EcopediaInput, FileName string) (Ecopedia, error) {
	newEcopedia := Ecopedia{}

	newEcopedia.Judul = ecopedia.Judul
	newEcopedia.Subjudul = ecopedia.Subjudul
	newEcopedia.Deskripsi = ecopedia.Deskripsi
	// newEcopedia.Gambar = ecopedia.Gambar
	newEcopedia.Gambar = FileName
	newEcopedia.Srcgambar = ecopedia.Srcgambar
	newEcopedia.Referensi = ecopedia.Referensi

	saveEcopedia, err := s.repository.Create(newEcopedia)

	if err != nil {
		return saveEcopedia, err
	}
	return saveEcopedia, nil
}
