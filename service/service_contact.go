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

type ServiceContact interface {
	SubmitContactSubmission(input dto.ContactSubmissionInput) (*entity.Contact, error)
	GetAllContactSubmission() ([]*entity.Contact, error)
	GetContactSubmissionById(slug string) (*entity.Contact, error)
	DeleteContactSubmission(slug string) (*entity.Contact, error)
}

type service_contact struct {
	repository repository.RepositoryContact
}

func NewServiceContact(repository repository.RepositoryContact) *service_contact {
	return &service_contact{repository}
}

func (s *service_contact) SubmitContactSubmission(input dto.ContactSubmissionInput) (*entity.Contact, error) {
	contact_submission := &entity.Contact{}

	contact_submission.Name = input.Name
	contact_submission.Email = input.Email
	contact_submission.Subject = input.Subject
	contact_submission.Message = input.Message

	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))

	slugTitle := strings.ToLower(input.Name)

	mySlug := slug.Make(slugTitle)

	randomNumber := seededRand.Intn(1000000) // Angka acak 0-999999

	contact_submission.Slug = fmt.Sprintf("%s-%d", mySlug, randomNumber)

	newContactSubmission, err := s.repository.Submit(contact_submission)
	if err != nil {
		return newContactSubmission, err
	}
	return newContactSubmission, nil
}

func (s *service_contact) GetAllContactSubmission() ([]*entity.Contact, error) {
	contact_submissions, err := s.repository.FindAll()
	if err != nil {
		return contact_submissions, err
	}
	return contact_submissions, nil
}

func (s *service_contact) GetContactSubmissionById(slug string) (*entity.Contact, error) {
	contact_submission, err := s.repository.FindBySlug(slug)

	if err != nil {
		return contact_submission, err
	}

	return contact_submission, nil
}

func (s *service_contact) DeleteContactSubmission(slug string) (*entity.Contact, error) {
	contact_submission, err := s.repository.FindBySlug(slug)
	if err != nil {
		return contact_submission, err
	}

	contact_submissionDel, err := s.repository.Delete(contact_submission)

	if err != nil {
		return contact_submissionDel, err
	}
	return contact_submissionDel, nil
}
