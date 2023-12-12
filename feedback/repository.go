package feedback

import (
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]Feedback, error)
	FindById(id int) (Feedback, error)
	FindBySlug(slug string) (Feedback, error)
	Create(feedback Feedback) (Feedback, error)
	Delete(feedback Feedback) (Feedback, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Delete(feedback Feedback) (Feedback, error) {
	err := r.db.Delete(&feedback).Error
	if err != nil {
		return feedback, err
	}
	return feedback, nil
}

func (r *repository) FindAll() ([]Feedback, error) {
	var feedbacks []Feedback
	err := r.db.Find(&feedbacks).Error
	if err != nil {
		return feedbacks, err
	}
	return feedbacks, nil
}

func (r *repository) FindById(id int) (Feedback, error) {
	var feedback Feedback
	err := r.db.Find(&feedback, id).Error
	if err != nil {
		return feedback, err
	}
	return feedback, nil
}

func (r *repository) FindBySlug(slug string) (Feedback, error) {
	var feedback Feedback

	err := r.db.Where("slug = ?", slug).Find(&feedback).Error

	if err != nil {
		return feedback, err
	}
	if feedback.Slug == "" {
		return feedback, errors.New("slug not found")
	}

	return feedback, nil

}

func (r *repository) Create(feedback Feedback) (Feedback, error) {
	err := r.db.Create(&feedback).Error
	if err != nil {
		return feedback, err
	}
	return feedback, nil
}
