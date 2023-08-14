package contact

import (
	"errors"
)

type Service interface {
	SubmitContactSubmission(input ContactSubmissionInput) (Contact, error)
	GetContactSubmissionById(ID int) (Contact, error)
	GetContactSubmissionByName(Name string) (Contact, error)
	GetContactSubmissionByEmail(Email string) (Contact, error)
	DeleteContactSubmission(ID int) (Contact, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) SubmitContactSubmission(input ContactSubmissionInput) (Contact, error) {
	contact_submission := Contact{}

	contact_submission.Name = input.Name
	contact_submission.Email = input.Email
	contact_submission.Subject = input.Subject
	contact_submission.Message = input.Message

	newContactSubmission, err := s.repository.Submit(contact_submission)
	if err != nil {
		return newContactSubmission, err
	}
	return newContactSubmission, nil
}

func (s *service) GetContactSubmissionById(ID int) (Contact, error) {
	contact_submission, err := s.repository.FindById(ID)

	if err != nil {
		return contact_submission, err
	}

	if contact_submission.ID == 0 {
		return contact_submission, errors.New("contact_submission Not Found With That ID")
	}

	return contact_submission, nil
}

func (s *service) GetContactSubmissionByName(Name string) (Contact, error) {
	contact_submission, err := s.repository.FIndByName(Name)

	if err != nil {
		return contact_submission, err
	}

	return contact_submission, nil
}

func (s *service) GetContactSubmissionByEmail(Email string) (Contact, error) {
	contact_submission, err := s.repository.FindByEmail(Email)

	if err != nil {
		return contact_submission, err
	}

	return contact_submission, nil
}

func (s *service) DeleteContactSubmission(ID int) (Contact, error) {
	contact_submission, err := s.repository.FindById(ID)
	if err != nil {
		return contact_submission, err
	}
	contact_submissionDel, err := s.repository.Delete(contact_submission)

	if err != nil {
		return contact_submissionDel, err
	}
	return contact_submissionDel, nil
}