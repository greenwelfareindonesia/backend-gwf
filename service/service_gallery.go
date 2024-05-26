package service

import (
	"fmt"
	"greenwelfare/dto"
	"greenwelfare/entity"
	"greenwelfare/repository"
	"math/rand"
	"strings"
	"time"

	"github.com/gosimple/slug"
)

type ServiceGallery interface {
	CreateGallery(input dto.InputGallery) (*entity.Gallery, error)
	CreateImageGallery(galleryID int, fileLocation string) error
	GetAllGallery(input int) ([]*entity.Gallery, error)
	GetOneGallery(slug string) (*entity.Gallery, error)
	UpdateGallery(slug string, input dto.InputGallery) (*entity.Gallery, error)
	DeleteGallery(ID int) (*entity.Gallery, error)
}

type service_gallery struct {
	repository repository.RepositoryGallery
}

func NewServiceGallery(repository repository.RepositoryGallery) *service_gallery {
	return &service_gallery{repository}
}

func (s *service_gallery) CreateImageGallery(galleryID int, fileLocation string) error {
	createImage := &entity.GalleryImages{}

	createImage.FileName = fileLocation
	createImage.GalleryID = galleryID

	err := s.repository.CreateImage(createImage)
	if err != nil {
		return err
	}
	return nil
}

func (s *service_gallery) CreateGallery(input dto.InputGallery) (*entity.Gallery, error) {
	addGalleryImage := &entity.Gallery{}

	// addGalleryImage.Image = fileLocation
	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))
	addGalleryImage.Alt = input.Alt

	slugTitle := strings.ToLower(input.Alt)

	mySlug := slug.Make(slugTitle)

	randomNumber := seededRand.Intn(1000000) // Angka acak 0-999999

	addGalleryImage.Slug = fmt.Sprintf("%s-%d", mySlug, randomNumber)

	//addGalleryImage.Likes = input.Likes

	newGalleryImage, err := s.repository.Create(addGalleryImage)
	if err != nil {
		return newGalleryImage, err
	}
	return newGalleryImage, nil
}

func (s *service_gallery) GetAllGallery(input int) ([]*entity.Gallery, error) {
	gallery, err := s.repository.FindAll()
	if err != nil {
		return gallery, err
	}
	return gallery, nil
}

func (s *service_gallery) GetOneGallery(slug string) (*entity.Gallery, error) {
	gallery, err := s.repository.FindBySlug(slug)
	if err != nil {
		return gallery, err
	}
	if gallery.ID == 0 {
		return gallery, err
	}
	return gallery, nil
}

func (s *service_gallery) UpdateGallery(slugs string, input dto.InputGallery) (*entity.Gallery, error) {
	addGalleryImage, err := s.repository.FindBySlug(slugs)
	if err != nil {
		return addGalleryImage, err
	}

	oldSlug := addGalleryImage.Slug

	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	slugTitle := strings.ToLower(addGalleryImage.Alt)
	mySlug := slug.Make(slugTitle)
	randomNumber := seededRand.Intn(1000000) // Angka acak 0-999999
	addGalleryImage.Slug = fmt.Sprintf("%s-%d", mySlug, randomNumber)

	// Ubah nilai slug kembali ke nilai slug lama untuk mencegah perubahan slug dalam database
	addGalleryImage.Slug = oldSlug

	addGalleryImage.Alt = input.Alt
	//addGalleryImage.Likes = input.Likes

	// Update the addGalleryImage in the repository
	newGalleryImage, err := s.repository.Update(addGalleryImage)
	if err != nil {
		return addGalleryImage, err
	}

	return newGalleryImage, nil
}

func (s *service_gallery) DeleteGallery(ID int) (*entity.Gallery, error) {
	galleryImage, err := s.repository.FindById(ID)
	if err != nil {
		return galleryImage, err
	}

	err = s.repository.DeleteImages(galleryImage.ID)
	if err != nil {
		return galleryImage, err
	}

	newWorkshop, err := s.repository.Delete(galleryImage)
	if err != nil {
		return newWorkshop, err
	}
	return newWorkshop, nil
}
