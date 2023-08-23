package ecopedia

type Service interface {
	GetAllEcopedia(input int) ([]Ecopedia, error)
	GetEcopediaByID(id int) (Ecopedia, error)
	CreateEcopedia(ecopedia EcopediaInput, FileName string) (Ecopedia, error)
	DeleteEcopedia(ID int) (Ecopedia, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) DeleteEcopedia(ID int) (Ecopedia, error) {
	ecopedias, err := s.repository.FindById(ID)
	if err != nil {
		return ecopedias, err
	}

	ecopedia, err := s.repository.DeleteEcopedia(ecopedias)
	if err != nil {
		return ecopedia, err
	}
	return ecopedia, nil
}

func (s *service) GetAllEcopedia(input int) ([]Ecopedia, error) {
	ecopedias, err := s.repository.FindAll()
	if err != nil {
		return ecopedias, err
	}
	return ecopedias, nil
}

func (s *service) GetEcopediaByID(id int) (Ecopedia, error) {
	ecopedias, err := s.repository.FindById(id)
	if err != nil {
		return ecopedias, err
	}
	return ecopedias, nil
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
