package gallery

import (
	"gorm.io/gorm"
)

type Repository interface {
	Create(gallery Gallery) (Gallery, error)
	FindAll() ([]Gallery, error)
	FindById(ID int) (Gallery, error)
	Update(gallery Gallery) (Gallery, error)
	Delete(gallery Gallery) (Gallery, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Create(gallery Gallery) (Gallery, error) {
	err := r.db.Create(&gallery).Error

	if err != nil {
		return gallery, err
	}
	return gallery, nil
}

func (r *repository) FindAll() ([]Gallery, error) {
	var gallery []Gallery

	err := r.db.Find(&gallery).Error

	if err != nil {
		return gallery, err
	}
	return gallery, nil
}

func (r *repository) FindById(ID int) (Gallery, error) {
	var gallery Gallery

	err := r.db.Where("id = ?", ID).Find(&gallery).Error

	if err != nil {
		return gallery, err
	}
	return gallery, nil
}

func (r *repository) Update(gallery Gallery) (Gallery, error) {
	err := r.db.Save(&gallery).Error

	if err != nil {
		return gallery, err
	}
	return gallery, nil
}

func (r *repository) Delete(gallery Gallery) (Gallery, error) {
	err := r.db.Delete(&gallery).Error

	if err != nil {
		return gallery, err
	}
	return gallery, nil
}
