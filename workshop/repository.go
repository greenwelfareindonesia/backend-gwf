package workshop

import (
	"gorm.io/gorm"
)

type Repository interface {
	Create(workshop Workshop) (Workshop, error)
	FindAll() ([]Workshop, error)
	FindById(ID int) (Workshop, error)
	Update(workshop Workshop) (Workshop, error)
	Delete(workshop Workshop) (Workshop, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Create(workshop Workshop) (Workshop, error) {
	err := r.db.Create(&workshop).Error

	if err != nil {
		return workshop, err
	}
	return workshop, nil
}

func (r *repository) FindAll() ([]Workshop, error) {
	var workshop []Workshop

	err := r.db.Find(&workshop).Error

	if err != nil {
		return workshop, err
	}
	return workshop, nil
}

func (r *repository) FindById(ID int) (Workshop, error) {
	var workshop Workshop

	err := r.db.Where("id = ?", ID).Find(&workshop).Error

	if err != nil {
		return workshop, err
	}
	return workshop, nil
}

func (r *repository) Update(workshop Workshop) (Workshop, error) {
	err := r.db.Save(&workshop).Error

	if err != nil {
		return workshop, err
	}
	return workshop, nil
}

func (r *repository) Delete(workshop Workshop) (Workshop, error) {
	err := r.db.Delete(&workshop).Error

	if err != nil {
		return workshop, err
	}
	return workshop, nil
}
