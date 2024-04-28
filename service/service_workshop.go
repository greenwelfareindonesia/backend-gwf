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

type ServiceWorkshop interface {
	CreateWorkshop(input dto.CreateWorkshop, fileLocation string) (*entity.Workshop, error)
	GetAllWorkshop(input int) ([]*entity.Workshop, error)
	GetOneWorkshop(slugs string) (*entity.Workshop, error)
	DeleteWorkshop(slugs string) (*entity.Workshop, error)
	UpdateWorkshop(slugs string, input dto.CreateWorkshop, fileLocation string) (*entity.Workshop, error)
}

type service_workshop struct {
	repository repository.RepositoryWorkshop
}

func NewServiceWorkshop(repository repository.RepositoryWorkshop) *service_workshop {
	return &service_workshop{repository}
}

func (s *service_workshop) CreateWorkshop(input dto.CreateWorkshop, fileLocation string) (*entity.Workshop, error) {
	createWorkshop := &entity.Workshop{}

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

func (s *service_workshop) GetAllWorkshop(input int) ([]*entity.Workshop, error) {
	workshop, err := s.repository.FindAll()
	if err != nil {
		return workshop, err
	}
	return workshop, nil
}

func (s *service_workshop) GetOneWorkshop(slugs string) (*entity.Workshop, error) {
	workshop, err := s.repository.FindBySlug(slugs)
	if err != nil {
		return workshop, err
	}

	return workshop, nil
}

func (s *service_workshop) UpdateWorkshop(slugs string, input dto.CreateWorkshop, fileLocation string) (*entity.Workshop, error) {
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

func (s *service_workshop) DeleteWorkshop(slugs string) (*entity.Workshop, error) {
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
