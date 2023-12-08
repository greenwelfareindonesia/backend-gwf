package gallery

type Service interface {
	CreateGallery(input InputGallery, fileLocation string) (Gallery, error)
	GetAllGallery(input int) ([]Gallery, error)
	GetOneGallery(ID int) (Gallery, error)
	UpdateGallery(getIdGallery InputGalleryID, input InputGallery, fileLocation string) (Gallery, error)
	DeleteGallery(ID int) (Gallery, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateGallery(input InputGallery, fileLocation string) (Gallery, error) {
	addGalleryImage := Gallery{}

	addGalleryImage.Image = fileLocation
	addGalleryImage.Alt = input.Alt
	//addGalleryImage.Likes = input.Likes

	newGalleryImage, err := s.repository.Create(addGalleryImage)
	if err != nil {
		return newGalleryImage, err
	}
	return newGalleryImage, nil
}

func (s *service) GetAllGallery(input int) ([]Gallery, error) {
	gallery, err := s.repository.FindAll()
	if err != nil {
		return gallery, err
	}
	return gallery, nil
}

func (s *service) GetOneGallery(ID int) (Gallery, error) {
	gallery, err := s.repository.FindById(ID)
	if err != nil {
		return gallery, err
	}
	if gallery.ID == 0 {
		return gallery, err
	}
	return gallery, nil
}

func (s *service) UpdateGallery(getIdGallery InputGalleryID, input InputGallery, fileLocation string) (Gallery, error) {
	addGalleryImage, err := s.repository.FindById(getIdGallery.ID)
	if err != nil {
		return addGalleryImage, err
	}

	// Update the addGalleryImage properties with the new values
	addGalleryImage.Image = fileLocation
	addGalleryImage.Alt = input.Alt
	//addGalleryImage.Likes = input.Likes

	// Update the addGalleryImage in the repository
	newGalleryImage, err := s.repository.Update(addGalleryImage)
	if err != nil {
		return addGalleryImage, err
	}

	return newGalleryImage, nil
}

func (s *service) DeleteGallery(ID int) (Gallery, error) {
	galleryImage, err := s.repository.FindById(ID)
	if err != nil {
		return galleryImage, err
	}

	newWorkshop, err := s.repository.Delete(galleryImage)
	if err != nil {
		return newWorkshop, err
	}
	return newWorkshop, nil
}
