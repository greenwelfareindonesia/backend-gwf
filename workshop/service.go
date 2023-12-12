package workshop

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/gosimple/slug"
)

type Service interface {
	CreateWorkshop(input CreateWorkshop, fileLocation string) (Workshop, error)
	GetAllWorkshop(input int) ([]Workshop, error)
	GetOneWorkshop(slugs string) (Workshop, error)
	DeleteWorkshop(slugs string) (Workshop, error)
	UpdateWorkshop(slugs string, input CreateWorkshop, fileLocation string) (Workshop, error)
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
	createWorkshop.Desc = input.Description
	createWorkshop.Date = input.Date
	createWorkshop.Url = input.Url
	createWorkshop.IsOpen = input.IsOpen

	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))

	slugTitle := strings.ToLower(input.Title)

	mySlug := slug.Make(slugTitle)

	randomNumber := seededRand.Intn(1000000) // Angka acak 0-999999

	createWorkshop.Slug = fmt.Sprintf("%s-%d", mySlug, randomNumber)

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

func (s *service) GetOneWorkshop(slugs string) (Workshop, error) {
	workshop, err := s.repository.FindBySlug(slugs)
	if err != nil {
		return workshop, err
	}

	return workshop, nil
}

func (s *service) UpdateWorkshop(slugs string, input CreateWorkshop, fileLocation string) (Workshop, error) {
	workshop, err := s.repository.FindBySlug(slugs)
	if err != nil {
		return workshop, err
	}

	oldSlug := workshop.Slug
	// Update the workshop properties with the new values
	workshop.Title = input.Title
	workshop.Image = fileLocation
	workshop.Desc = input.Description
	workshop.Date = input.Date
	workshop.Url = input.Url
	workshop.IsOpen = input.IsOpen

	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	slugTitle := strings.ToLower(workshop.Title)
	mySlug := slug.Make(slugTitle)
	randomNumber := seededRand.Intn(1000000) // Angka acak 0-999999
	workshop.Slug = fmt.Sprintf("%s-%d", mySlug, randomNumber)

	// Ubah nilai slug kembali ke nilai slug lama untuk mencegah perubahan slug dalam database
	workshop.Slug = oldSlug

	// Update the workshop in the repository
	newWorkshop, err := s.repository.Update(workshop)
	if err != nil {
		return workshop, err
	}

	return newWorkshop, nil
}

func (s *service) DeleteWorkshop(slugs string) (Workshop, error) {
	workshop, err := s.repository.FindBySlug(slugs)
	if err != nil {
		return workshop, err
	}

	newWorkshop, err := s.repository.Delete(workshop)
	if err != nil {
		return newWorkshop, err
	}
	return newWorkshop, nil
}
