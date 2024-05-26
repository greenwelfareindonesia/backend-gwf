package repository

import (
	"errors"
	"greenwelfare/entity"

	"gorm.io/gorm"
)

type RepositoryVeganguide interface {
	FindAll() ([]*entity.Veganguide, error)
	FindById(id int) (*entity.Veganguide, error)
	FindBySlug(slug string) (*entity.Veganguide, error)
	Create(veganguide *entity.Veganguide) (*entity.Veganguide, error)
	DeleteVeganguide(veganguide *entity.Veganguide) (*entity.Veganguide, error)
	Update(veganguide *entity.Veganguide) (*entity.Veganguide, error)
}

type repository_veganguide struct {
	db *gorm.DB
}

func NewRepositoryVeganguide(db *gorm.DB) *repository_veganguide {
	return &repository_veganguide{db}
}

func (r *repository_veganguide) DeleteVeganguide(veganguide *entity.Veganguide) (*entity.Veganguide, error) {
	err := r.db.Delete(&veganguide).Error
	if err != nil {
		return veganguide, err
	}
	return veganguide, nil
}

func (r *repository_veganguide) FindAll() ([]*entity.Veganguide, error) {
	var veganguides []*entity.Veganguide
	err := r.db.Find(&veganguides).Error
	if err != nil {
		return veganguides, err
	}
	return veganguides, nil
}

func (r *repository_veganguide) FindById(id int) (*entity.Veganguide, error) {
	var veganguide *entity.Veganguide
	err := r.db.Where("id = ?", id).Find(&veganguide).Error
	if err != nil {
		return veganguide, err
	}
	return veganguide, nil
}

func (r *repository_veganguide) FindBySlug(slug string) (*entity.Veganguide, error) {
	var veganguide *entity.Veganguide

	err := r.db.Where("slug = ?", slug).Find(&veganguide).Error

	if err != nil {
		return veganguide, err
	}
	if veganguide.Slug == "" {
		return veganguide, errors.New("slug not found")
	}

	return veganguide, nil

}

func (r *repository_veganguide) Create(veganguide *entity.Veganguide) (*entity.Veganguide, error) {
	err := r.db.Create(&veganguide).Error
	if err != nil {
		return veganguide, err
	}
	return veganguide, nil
}

func (r *repository_veganguide) Update(veganguide *entity.Veganguide) (*entity.Veganguide, error) {
	err := r.db.Save(&veganguide).Error
	if err != nil {
		return veganguide, err
	}

	return veganguide, nil
}
