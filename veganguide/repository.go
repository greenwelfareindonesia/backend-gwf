package veganguide

import (
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]Veganguide, error)
	FindById(id int) (Veganguide, error)
	FindBySlug(slug string) (Veganguide, error)
	Create(veganguide Veganguide) (Veganguide, error)
	DeleteVeganguide(veganguide Veganguide) (Veganguide, error)
	Update(veganguide Veganguide) (Veganguide, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) DeleteVeganguide(veganguide Veganguide) (Veganguide, error) {
	err := r.db.Delete(&veganguide).Error
	if err != nil {
		return veganguide, err
	}
	return veganguide, nil
}

func (r *repository) FindAll() ([]Veganguide, error) {
	var veganguides []Veganguide
	err := r.db.Find(&veganguides).Error
	if err != nil {
		return veganguides, err
	}
	return veganguides, nil
}

func (r *repository) FindById(id int) (Veganguide, error) {
	var veganguide Veganguide
	err := r.db.Where("id = ?", id).Find(&veganguide).Error
	if err != nil {
		return veganguide, err
	}
	return veganguide, nil
}

func (r *repository) FindBySlug(slug string) (Veganguide, error) {
	var veganguide Veganguide

	err := r.db.Where("slug = ?", slug).Find(&veganguide).Error

	if err != nil {
		return veganguide, err
	}
	if veganguide.Slug == "" {
		return veganguide, errors.New("slug not found")
	}

	return veganguide, nil

}

func (r *repository) Create(veganguide Veganguide) (Veganguide, error) {
	err := r.db.Create(&veganguide).Error
	if err != nil {
		return veganguide, err
	}
	return veganguide, nil
}

func (r *repository) Update(veganguide Veganguide) (Veganguide, error) {
	err := r.db.Save(&veganguide).Error
	if err != nil {
		return veganguide, err
	}

	return veganguide, nil
}
