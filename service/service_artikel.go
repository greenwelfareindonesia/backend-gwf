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

type ServiceArtikel interface {
	CreateArtikel(input dto.CreateArtikel) (*entity.Artikel, error)
	GetAllArtikel(input int) ([]*entity.Artikel, error)
	DeleteArtikel(slugs string) (*entity.Artikel, error)
	GetOneArtikel(slugs string) (*entity.Artikel, error)
	UpdateArtikel(input dto.CreateArtikel, slugs string) (*entity.Artikel, error)
}

type service_artikel struct {
	repository repository.RepositoryArtikel
}

func NewServiceArtikel(repository repository.RepositoryArtikel) *service_artikel {
	return &service_artikel{repository}
}

func (s *service_artikel) GetOneArtikel(slugs string) (*entity.Artikel, error) {
	berita, err := s.repository.FindBySlug(slugs)
	if err != nil {
		return berita, err
	}
	
	return berita, nil
}

func (s *service_artikel) UpdateArtikel(input dto.CreateArtikel, slugs string) (*entity.Artikel, error) {
	artikel, err := s.repository.FindBySlug(slugs)
	if err != nil {
		return artikel, err
	}

	oldSlug := artikel.Slug

	artikel.FullName = input.FullName
	artikel.Email = input.Email
	artikel.Topic = input.Topic
	artikel.ArticleMessage = input.ArticleMessage

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

func (s *service_artikel) DeleteArtikel(slugs string) (*entity.Artikel, error) {
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

func (s *service_artikel) GetAllArtikel(input int) ([]*entity.Artikel, error) {
	berita, err := s.repository.FindAll()
	if err != nil {
		return berita, err
	}
	return berita, nil
}

func (s *service_artikel) CreateArtikel(input dto.CreateArtikel) (*entity.Artikel, error) {
	createBerita := &entity.Artikel{}

	createBerita.FullName = input.FullName
	createBerita.Email = input.Email
	createBerita.Topic = input.Topic
	createBerita.ArticleMessage = input.ArticleMessage

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
