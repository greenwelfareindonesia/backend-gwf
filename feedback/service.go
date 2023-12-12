package feedback

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/gosimple/slug"
)

type Service interface {
	GetAllFeedback(input int) ([]Feedback, error)
	GetFeedbackBySlug(slugs string) (Feedback, error)
	CreateFeedback(feedback FeedbackInput) (Feedback, error)
	DeleteFeedback(slugs string) (Feedback, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) DeleteFeedback(slugs string) (Feedback, error) {
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

func (s *service) GetAllFeedback(input int) ([]Feedback, error) {
	feedbacks, err := s.repository.FindAll()
	if err != nil {
		return feedbacks, err
	}
	return feedbacks, nil
}

func (s *service) GetFeedbackBySlug(slugs string) (Feedback, error) {
	feedbacks, err := s.repository.FindBySlug(slugs)
	if err != nil {
		return feedbacks, err
	}

	return feedbacks, nil
}

func (s *service) CreateFeedback(feedback FeedbackInput) (Feedback, error) {
	newFeedback := Feedback{}

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
