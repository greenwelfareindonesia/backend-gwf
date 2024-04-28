package repository

import (
	"errors"
	"greenwelfare/entity"

	"gorm.io/gorm"
)

type RepositoryWorkshop interface {
	Create(workshop *entity.Workshop) (*entity.Workshop, error)
	FindAll() ([]*entity.Workshop, error)
	FindById(id int) (*entity.Workshop, error)
	FindBySlug(slug string) (*entity.Workshop, error)
	Update(workshop *entity.Workshop) (*entity.Workshop, error)
	Delete(workshop *entity.Workshop) (*entity.Workshop, error)
}

type repository_workshop struct {
	db *gorm.DB
}

func NewRepositoryWorkshop(db *gorm.DB) *repository_workshop {
	return &repository_workshop{db}
}

func (r *repository_workshop) Create(workshop *entity.Workshop) (*entity.Workshop, error) {
	err := r.db.Create(&workshop).Error

	if err != nil {
		return workshop, err
	}
	return workshop, nil
}

func (r *repository_workshop) FindAll() ([]*entity.Workshop, error) {
	var workshop []*entity.Workshop

	err := r.db.Find(&workshop).Error

	if err != nil {
		return workshop, err
	}
	return workshop, nil
}

func (r *repository_workshop) FindById(ID int) (*entity.Workshop, error) {
	var workshop *entity.Workshop

	err := r.db.Where("id = ?", ID).Find(&workshop).Error

	if err != nil {
		return workshop, err
	}
	return workshop, nil
}

func (r *repository_workshop) FindBySlug(slug string) (*entity.Workshop, error) {
	var workshop *entity.Workshop

	err := r.db.Where("slug = ?", slug).Find(&workshop).Error

	if err != nil {
		return workshop, err
	}
	if workshop.Slug == "" {
		return workshop, errors.New("slug not found")
	}

	return workshop, nil

}

func (r *repository_workshop) Update(workshop *entity.Workshop) (*entity.Workshop, error) {
	err := r.db.Save(&workshop).Error

	if err != nil {
		return workshop, err
	}
	return workshop, nil
}

func (r *repository_workshop) Delete(workshop *entity.Workshop) (*entity.Workshop, error) {
	err := r.db.Delete(&workshop).Error

	if err != nil {
		return workshop, err
	}
	return workshop, nil
}
