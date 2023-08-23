package event

type Service interface {
	CreateEvent(input CreateEvents, fileLocation string) (Event, error)
	GetAllEvent(input int) ([]Event, error)
	DeleteEvent(ID int) (Event, error)
	GetOneEvent(ID int) (Event, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetOneEvent(ID int) (Event, error) {
	berita, err := s.repository.FindById(ID)
	if err != nil {
		return berita, err
	}
	return berita, nil
}

func (s *service) DeleteEvent(ID int) (Event, error) {
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

func (s *service) GetAllEvent(input int) ([]Event, error) {
	berita, err := s.repository.FindAll()
	if err != nil {
		return berita, err
	}
	return berita, nil
}

func (s *service) CreateEvent(input CreateEvents, fileLocation string) (Event, error) {
	createBerita := Event{}

	createBerita.Judul = input.Judul
	createBerita.EventMessage = input.EventMessage
	createBerita.FileName = fileLocation

	newBerita, err := s.repository.Save(createBerita)
	if err != nil {
		return newBerita, err
	}
	return newBerita, nil
}