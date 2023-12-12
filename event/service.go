package event

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/gosimple/slug"
)

type Service interface {
	CreateEvent(input CreateEvents, fileLocation string) (Event, error)
	GetAllEvent(input int) ([]Event, error)
	DeleteEvent(slug string) (Event, error)
	GetOneEvent(slug string) (Event, error)
	UpdateEvent(slugs string, input CreateEvents, FileLocation string) (Event, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetOneEvent(slug string) (Event, error) {
	berita, err := s.repository.FindBySlug(slug)
	if err != nil {
		return berita, err
	}
	if berita.ID == 0 {
		return berita, err
	}
	return berita, nil
}

func (s *service) DeleteEvent(slug string) (Event, error) {
	berita, err := s.repository.FindBySlug(slug)
	if err != nil {
		return berita, err
	}

	newBerita, err := s.repository.Delete(berita)
	if err != nil {
		return newBerita, err
	}
	return newBerita, nil
}

func (s *service) GetAllEvent(input int) ([]Event, error) {
	berita, err := s.repository.FindAll()
	if err != nil {
		return berita, err
	}
	return berita, nil
}

func (s *service) CreateEvent(input CreateEvents, fileLocation string) (Event, error) {
	createBerita := Event{}

	createBerita.Title = input.Title
	createBerita.EventMessage = input.EventMessage
	createBerita.FileName = fileLocation

	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))

	slugTitle := strings.ToLower(input.Title)

	mySlug := slug.Make(slugTitle)

	randomNumber := seededRand.Intn(1000000) // Angka acak 0-999999

	createBerita.Slug = fmt.Sprintf("%s-%d", mySlug, randomNumber)

	newBerita, err := s.repository.Save(createBerita)
	if err != nil {
		return newBerita, err
	}
	return newBerita, nil
}

func (s *service) UpdateEvent(slugs string, input CreateEvents, FileLocation string) (Event, error) {
	event, err := s.repository.FindBySlug(slugs)
	if err != nil {
		return event, err
	}

	oldSlug := event.Slug
	event.Title = input.Title
	event.EventMessage = input.EventMessage
	event.FileName = FileLocation

	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
    slugTitle := strings.ToLower(event.Title)
    mySlug := slug.Make(slugTitle)
    randomNumber := seededRand.Intn(1000000) // Angka acak 0-999999
    event.Slug = fmt.Sprintf("%s-%d", mySlug, randomNumber)

    // Ubah nilai slug kembali ke nilai slug lama untuk mencegah perubahan slug dalam database
    event.Slug = oldSlug

	newEvent, err := s.repository.Update(event)
	if err != nil {
		return newEvent, err
	}
	return newEvent, nil

}
