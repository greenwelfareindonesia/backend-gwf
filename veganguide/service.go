package veganguide

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/gosimple/slug"
)

type Service interface {
	GetAllVeganguide(input int) ([]Veganguide, error)
	GetOneVeganguide(slugs string) (Veganguide, error)
	CreateVeganguide(veganguide VeganguideInput, FileName string) (Veganguide, error)
	DeleteVeganguide(ID int) (Veganguide, error)
	UpdateVeganguide(input VeganguideInput, slugs string, FileName string) (Veganguide, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) DeleteVeganguide(ID int) (Veganguide, error) {
	veganguides, err := s.repository.FindById(ID)
	if err != nil {
		return veganguides, err
	}

	veganguide, err := s.repository.DeleteVeganguide(veganguides)
	if err != nil {
		return veganguide, err
	}
	return veganguide, nil
}

func (s *service) GetAllVeganguide(input int) ([]Veganguide, error) {
	veganguides, err := s.repository.FindAll()
	if err != nil {
		return veganguides, err
	}
	return veganguides, nil
}

func (s *service) GetOneVeganguide(slugs string) (Veganguide, error) {
	veganguides, err := s.repository.FindBySlug(slugs)
	if err != nil {
		return veganguides, err
	}
	if veganguides.ID == 0 {
		return veganguides, err
	}
	return veganguides, nil
}

func (s *service) CreateVeganguide(input VeganguideInput, FileName string) (Veganguide, error) {
	newVeganguide := Veganguide{}

	newVeganguide.Judul = input.Judul
	newVeganguide.Deskripsi = input.Deskripsi
	newVeganguide.Body = input.Body
	newVeganguide.Gambar = FileName

	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))

	slugTitle := strings.ToLower(input.Judul)
	mySlug := slug.Make(slugTitle)
	randomNumber := seededRand.Intn(1000000)

	newVeganguide.Slug = fmt.Sprintf("%s-%d", mySlug, randomNumber)

	saveVeganGuide, err := s.repository.Create(newVeganguide)

	if err != nil {
		return saveVeganGuide, err
	}
	return saveVeganGuide, nil
}

func (s *service) UpdateVeganguide(input VeganguideInput, slugs string, FileName string) (Veganguide, error) {
	veganguide, err := s.repository.FindBySlug(slugs)
	if err != nil {
		return veganguide, err
	}

	oldSlug := veganguide.Slug

	veganguide.Judul = input.Judul
	veganguide.Deskripsi = input.Deskripsi
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
