package artikel

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/gosimple/slug"
)

type Service interface {
	CreateArtikel(input CreateArtikel) (Artikel, error)
	GetAllArtikel(input int) ([]Artikel, error)
	DeleteArtikel(slugs string) (Artikel, error)
	GetOneArtikel(slugs string) (Artikel, error)
	UpdateArtikel(input CreateArtikel, slugs string) (Artikel, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetOneArtikel(slugs string) (Artikel, error) {
	berita, err := s.repository.FindBySlug(slugs)
	if err != nil {
		return berita, err
	}
	if berita.ID == 0 {
		return berita, err
	}
	return berita, nil
}

func (s *service) UpdateArtikel(input CreateArtikel, slugs string) (Artikel, error) {
	artikel, err := s.repository.FindBySlug(slugs)
	if err != nil {
		return artikel, err
	}

	oldSlug := artikel.Slug


	artikel.FullName = input.FullName
	artikel.Email = input.Email
	artikel.Topic = input.Topic
	artikel.ArtikelMessage = input.ArtikelMessage

	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
    slugTitle := strings.ToLower(artikel.FullName)
    mySlug := slug.Make(slugTitle)
    randomNumber := seededRand.Intn(1000000) // Angka acak 0-999999
    artikel.Slug = fmt.Sprintf("%s-%d", mySlug, randomNumber)

    // Ubah nilai slug kembali ke nilai slug lama untuk mencegah perubahan slug dalam database
    artikel.Slug = oldSlug

	newArtikel, err := s.repository.Update(artikel)
	if err != nil {
		return newArtikel, err
	}
	return newArtikel, nil
}

func (s *service) DeleteArtikel(slugs string) (Artikel, error) {
	berita, err := s.repository.FindBySlug(slugs)
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

func (s *service) CreateArtikel(input CreateArtikel) (Artikel, error) {
	createBerita := Artikel{}

	createBerita.FullName = input.FullName
	createBerita.Email = input.Email
	createBerita.Topic = input.Topic
	createBerita.ArtikelMessage = input.ArtikelMessage

	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))

	slugTitle := strings.ToLower(input.Topic)

	mySlug := slug.Make(slugTitle)

	randomNumber := seededRand.Intn(1000000) // Angka acak 0-999999

	createBerita.Slug = fmt.Sprintf("%s-%d", mySlug, randomNumber)

	newBerita, err := s.repository.Save(createBerita)
	if err != nil {
		return newBerita, err
	}
	return newBerita, nil
}
