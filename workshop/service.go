package workshop

type Service interface {
	CreateWorkshop(input CreateWorkshop, fileLocation string) (Workshop, error)
	GetAllWorkshop(input int) ([]Workshop, error)
	GetOneWorkshop(ID int) (Workshop, error)
	DeleteWorkshop(ID int) (Workshop, error)
	UpdateWorkshop(getIdWorkshop GetWorkshop, input CreateWorkshop, fileLocation string) (Workshop, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateWorkshop(input CreateWorkshop, fileLocation string) (Workshop, error) {
	createWorkshop := Workshop{}

	createWorkshop.Title = input.Title
	createWorkshop.Image = fileLocation
	createWorkshop.Desc = input.Desc
	createWorkshop.Date = input.Date
	createWorkshop.Url = input.Url
	createWorkshop.IsOpen = input.IsOpen

	newWorkshop, err := s.repository.Create(createWorkshop)
	if err != nil {
		return newWorkshop, err
	}
	return newWorkshop, nil
}

func (s *service) GetAllWorkshop(input int) ([]Workshop, error) {
	workshop, err := s.repository.FindAll()
	if err != nil {
		return workshop, err
	}
	return workshop, nil
}

func (s *service) GetOneWorkshop(ID int) (Workshop, error) {
	workshop, err := s.repository.FindById(ID)
	if err != nil {
		return workshop, err
	}
	return workshop, nil
}

func (s *service) UpdateWorkshop(getIdWorkshop GetWorkshop, input CreateWorkshop, fileLocation string) (Workshop, error) {
	workshop, err := s.repository.FindById(getIdWorkshop.ID)
	if err != nil {
		return workshop, err
	}

	// Update the workshop properties with the new values
	workshop.Title = input.Title
	workshop.Image = fileLocation
	workshop.Desc = input.Desc
	workshop.Date = input.Date
	workshop.Url = input.Url
	workshop.IsOpen = input.IsOpen

	// Update the workshop in the repository
	newWorkshop, err := s.repository.Update(workshop)
	if err != nil {
		return workshop, err
	}

	return newWorkshop, nil
}

func (s *service) DeleteWorkshop(ID int) (Workshop, error) {
	workshop, err := s.repository.FindById(ID)
	if err != nil {
		return workshop, err
	}

	newWorkshop, err := s.repository.Delete(workshop)
	if err != nil {
		return newWorkshop, err
	}
	return newWorkshop, nil
}
