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

type ServiceFeedback interface {
	GetAllFeedback(input int) ([]*entity.Feedback, error)
	GetFeedbackBySlug(slugs string) (*entity.Feedback, error)
	CreateFeedback(input dto.FeedbackInput) (*entity.Feedback, error)
	DeleteFeedback(slugs string) (*entity.Feedback, error)
}

type service_feedback struct {
	repository repository.RepositoryFeedback
}

func NewServiceFeedback(repository repository.RepositoryFeedback) *service_feedback {
	return &service_feedback{repository}
}

func (s *service_feedback) DeleteFeedback(slugs string) (*entity.Feedback, error) {
	feedbacks, err := s.repository.FindBySlug(slugs)
	if err != nil {
		return feedbacks, err
	}

	feedback, err := s.repository.Delete(feedbacks)
	if err != nil {
		return feedback, err
	}
	return feedback, nil
}

func (s *service_feedback) GetAllFeedback(input int) ([]*entity.Feedback, error) {
	feedbacks, err := s.repository.FindAll()
	if err != nil {
		return feedbacks, err
	}
	return feedbacks, nil
}

func (s *service_feedback) GetFeedbackBySlug(slugs string) (*entity.Feedback, error) {
	feedbacks, err := s.repository.FindBySlug(slugs)
	if err != nil {
		return feedbacks, err
	}

	return feedbacks, nil
}

func (s *service_feedback) CreateFeedback(feedback dto.FeedbackInput) (*entity.Feedback, error) {
	newFeedback := &entity.Feedback{}

	newFeedback.Email = feedback.Email
	newFeedback.Text = feedback.Text

	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))

	slugTitle := strings.ToLower(feedback.Email)

	mySlug := slug.Make(slugTitle)

	randomNumber := seededRand.Intn(1000000) // Angka acak 0-999999

	newFeedback.Slug = fmt.Sprintf("%s-%d", mySlug, randomNumber)

	saveFeedback, err := s.repository.Create(newFeedback)

	if err != nil {
		return saveFeedback, err
	}
	return saveFeedback, nil
}
