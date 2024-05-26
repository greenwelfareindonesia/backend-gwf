package repository

import (
	"greenwelfare/entity"
	"errors"
	"gorm.io/gorm"
)

type RepositoryFeedback interface {
	FindAll() ([]*entity.Feedback, error)
	FindById(id int) (*entity.Feedback, error)
	FindBySlug(slug string) (*entity.Feedback, error)
	Create(feedback *entity.Feedback) (*entity.Feedback, error)
	Delete(feedback *entity.Feedback) (*entity.Feedback, error)
}

type repository_feedback struct {
	db *gorm.DB
}

func NewRepositoryFeedback(db *gorm.DB) *repository_feedback {
	return &repository_feedback{db}
}

func (r *repository_feedback) Delete(feedback *entity.Feedback) (*entity.Feedback, error) {
	err := r.db.Delete(&feedback).Error
	if err != nil {
		return feedback, err
	}
	return feedback, nil
}

func (r *repository_feedback) FindAll() ([]*entity.Feedback, error) {
	var feedbacks []*entity.Feedback
	err := r.db.Find(&feedbacks).Error
	if err != nil {
		return feedbacks, err
	}
	return feedbacks, nil
}

func (r *repository_feedback) FindById(id int) (*entity.Feedback, error) {
	var feedback *entity.Feedback
	err := r.db.Find(&feedback, id).Error
	if err != nil {
		return feedback, err
	}
	return feedback, nil
}

func (r *repository_feedback) FindBySlug(slug string) (*entity.Feedback, error) {
	var feedback *entity.Feedback

	err := r.db.Where("slug = ?", slug).Find(&feedback).Error

	if err != nil {
		return feedback, err
	}
	if feedback.Slug == "" {
		return feedback, errors.New("slug not found")
	}

	return feedback, nil

}

func (r *repository_feedback) Create(feedback *entity.Feedback) (*entity.Feedback, error) {
	err := r.db.Create(&feedback).Error
	if err != nil {
		return feedback, err
	}
	return feedback, nil
}
