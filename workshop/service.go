package workshop

type Service interface {
	CreateWorkshop(input CreateWorkshop, fileLocation string) (Workshop, error) // belum ngerti input gambar
	GetAllWorkshop(input int) ([]Workshop, error)
	GetOneWorkshop(ID int) (Workshop, error)
	UpdateWorkshop(ID int, input UpdateWorkshop) (Workshop, error)
	DeleteWorkshop(ID int) (Workshop, error)
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
	// createWorkshop.Image = input.Image // belum ngerti input gambar
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

func (s *service) UpdateWorkshop(ID int, input UpdateWorkshop) (Workshop, error) {
	existingWorkshop, err := s.repository.FindById(ID)
	if err != nil {
		return existingWorkshop, err
	}

	// Update the fields of the existing workshop with the new values
	if input.Title != "" {
		existingWorkshop.Title = input.Title
	}
	if input.Image != "" {
		existingWorkshop.Image = input.Image
	}
	if input.Desc != "" {
		existingWorkshop.Desc = input.Desc
	}
	// if !input.Date.IsZero() {
	// 	existingWorkshop.Date = input.Date
	// }
	if input.Date != "" {
		existingWorkshop.Date = input.Date
	}
	if input.Url != "" {
		existingWorkshop.Url = input.Url
	}
	if input.IsOpen != existingWorkshop.IsOpen {
		existingWorkshop.IsOpen = input.IsOpen
	}

	// Save the updated workshop to the repository
	updatedWorkshop, err := s.repository.Update(existingWorkshop)
	if err != nil {
		return updatedWorkshop, err
	}

	return updatedWorkshop, nil
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
