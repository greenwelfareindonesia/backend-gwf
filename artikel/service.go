package artikel

type Service interface {
	CreateArtikel(input CreateArtikel, fileLocation string) (Artikel, error)
	GetAllArtikel(input int) ([]Artikel, error)
	DeleteArtikel(ID int) (Artikel, error)
	GetOneArtikel(ID int) (Artikel, error)
	UpdateArtikel(input CreateArtikel, getIdArtikel GetArtikel) (Artikel, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetOneArtikel(ID int) (Artikel, error) {
	berita, err := s.repository.FindById(ID)
	if err != nil {
		return berita, err
	}
	return berita, nil
}

func (s *service) UpdateArtikel(input CreateArtikel, getIdArtikel GetArtikel) (Artikel, error) {
	artikel, err := s.repository.FindById(getIdArtikel.ID)
	if err != nil {
		return artikel, err
	}
	artikel.FullName = input.FullName
	artikel.Email = input.Email
	artikel.Topic = input.Topic
	artikel.ArtikelMessage = input.ArtikelMessage

	newArtikel, err := s.repository.Update(artikel)
	if err != nil {
		return newArtikel, err
	}
	return newArtikel, nil
}

func (s *service) DeleteArtikel(ID int) (Artikel, error) {
	berita, err := s.repository.FindById(ID)
	if err != nil {
		return berita, err
	}

	newBerita, err := s.repository.Delete(berita)
	if err != nil {
		return newBerita, err
	}
	return newBerita, nil
}

func (s *service) GetAllArtikel(input int) ([]Artikel, error) {
	berita, err := s.repository.FindAll()
	if err != nil {
		return berita, err
	}
	return berita, nil
}

func (s *service) CreateArtikel(input CreateArtikel, fileLocation string) (Artikel, error) {
	createBerita := Artikel{}

	createBerita.FullName = input.FullName
	createBerita.Email = input.Email
	createBerita.Topic = input.Topic
	createBerita.ArtikelMessage = input.ArtikelMessage

	newBerita, err := s.repository.Save(createBerita)
	if err != nil {
		return newBerita, err
	}
	return newBerita, nil
}
