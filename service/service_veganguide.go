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

type ServiceVeganguide interface {
	GetAllVeganguide(input int) ([]*entity.Veganguide, error)
	GetOneVeganguide(slugs string) (*entity.Veganguide, error)
	CreateVeganguide(veganguide dto.VeganguideInput, FileName string) (*entity.Veganguide, error)
	DeleteVeganguide(slug string) (*entity.Veganguide, error)
	UpdateVeganguide(input dto.VeganguideInput, slugs string, FileName string) (*entity.Veganguide, error)
}

type service_veganguide struct {
	repository repository.RepositoryVeganguide
}

func NewServiceVeganguide(repository repository.RepositoryVeganguide) *service_veganguide {
	return &service_veganguide{repository}
}

func (s *service_veganguide) DeleteVeganguide(slug string) (*entity.Veganguide, error) {
	veganguides, err := s.repository.FindBySlug(slug)
	if err != nil {
		return veganguides, err
	}

	veganguide, err := s.repository.DeleteVeganguide(veganguides)
	if err != nil {
		return veganguide, err
	}
	return veganguide, nil
}

func (s *service_veganguide) GetAllVeganguide(input int) ([]*entity.Veganguide, error) {
	veganguides, err := s.repository.FindAll()
	if err != nil {
		return veganguides, err
	}
	return veganguides, nil
}

func (s *service_veganguide) GetOneVeganguide(slugs string) (*entity.Veganguide, error) {
	veganguides, err := s.repository.FindBySlug(slugs)
	if err != nil {
		return veganguides, err
	}
	if veganguides.ID == 0 {
		return veganguides, err
	}
	return veganguides, nil
}

func (s *service_veganguide) CreateVeganguide(input dto.VeganguideInput, FileName string) (*entity.Veganguide, error) {
	newVeganguide := &entity.Veganguide{}

	newVeganguide.Judul = input.Title
	newVeganguide.Deskripsi = input.Description
	newVeganguide.Body = input.Body
	newVeganguide.Gambar = FileName

	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))

	slugTitle := strings.ToLower(input.Title)
	mySlug := slug.Make(slugTitle)
	randomNumber := seededRand.Intn(1000000)

	newVeganguide.Slug = fmt.Sprintf("%s-%d", mySlug, randomNumber)

	saveVeganGuide, err := s.repository.Create(newVeganguide)

	if err != nil {
		return saveVeganGuide, err
	}
	return saveVeganGuide, nil
}

func (s *service_veganguide) UpdateVeganguide(input dto.VeganguideInput, slugs string, FileName string) (*entity.Veganguide, error) {
	veganguide, err := s.repository.FindBySlug(slugs)
	if err != nil {
		return veganguide, err
	}

	oldSlug := veganguide.Slug

	veganguide.Judul = input.Title
	veganguide.Deskripsi = input.Description
	veganguide.Body = input.Body
	veganguide.Gambar = FileName

	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	slugTitle := strings.ToLower(veganguide.Judul)
	mySlug := slug.Make(slugTitle)
	randomNumber := seededRand.Intn(1000000) // Angka acak 0-999999
	veganguide.Slug = fmt.Sprintf("%s-%d", mySlug, randomNumber)

	// Ubah nilai slug kembali ke nilai slug lama untuk mencegah perubahan slug dalam database
	veganguide.Slug = oldSlug

	newVeganguide, err := s.repository.Update(veganguide)
	if err != nil {
		return newVeganguide, err
	}
	return newVeganguide, nil
}
