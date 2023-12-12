package feedback

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Feedback, error)
	FindById(id int) (Feedback, error)
	Create(feedback Feedback) (Feedback, error)
	DeleteFeedback(feedback Feedback) (Feedback, error)
	Update(feedback Feedback) (Feedback, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Update(feedback Feedback) (Feedback, error) {
	err := r.db.Save(&feedback).Error
	if err != nil {
		return feedback, err
	}

	return feedback, nil

}

func (r *repository) DeleteFeedback(feedback Feedback) (Feedback, error) {
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

func (r *repository) Create(feedback Feedback) (Feedback, error) {
	err := r.db.Create(&feedback).Error
	if err != nil {
		return feedback, err
	}
	return feedback, nil
}
