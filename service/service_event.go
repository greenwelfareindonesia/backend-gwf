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

type ServiceEvent interface {
	CreateEvent(input dto.CreateEvents, fileLocation string) (*entity.Event, error)
	GetAllEvent(input int) ([]*entity.Event, error)
	DeleteEvent(slug string) (*entity.Event, error)
	GetOneEvent(slug string) (*entity.Event, error)
	UpdateEvent(slugs string, input dto.CreateEvents, FileLocation string) (*entity.Event, error)
}

type service_event struct {
	repository repository.RepositoryEvent
}

func NewServiceEvent(repository repository.RepositoryEvent) *service_event {
	return &service_event{repository}
}

func (s *service_event) GetOneEvent(slug string) (*entity.Event, error) {
	berita, err := s.repository.FindBySlug(slug)
	if err != nil {
		return berita, err
	}
	if berita.ID == 0 {
		return berita, err
	}
	return berita, nil
}

func (s *service_event) DeleteEvent(slug string) (*entity.Event, error) {
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

func (s *service_event) GetAllEvent(input int) ([]*entity.Event, error) {
	berita, err := s.repository.FindAll()
	if err != nil {
		return berita, err
	}
	return berita, nil
}

func (s *service_event) CreateEvent(input dto.CreateEvents, fileLocation string) (*entity.Event, error) {
	createBerita := &entity.Event{}

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

func (s *service_event) UpdateEvent(slugs string, input dto.CreateEvents, FileLocation string) (*entity.Event, error) {
	event, err := s.repository.FindBySlug(slugs)
	if err != nil {
		return event, err
	}

	event.Title = input.Title
	event.EventMessage = input.EventMessage
	event.FileName = FileLocation
	oldSlug := event.Slug


	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
    slugTitle := strings.ToLower(input.Title)
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
