package ecopedia

import "gorm.io/gorm"

type Repository interface {
	Create(ecopedia Ecopedia) (Ecopedia, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Create(ecopedia Ecopedia) (Ecopedia, error) {
	err := r.db.Create(&ecopedia).Error
	if err != nil {
		return ecopedia, err
	}
	return ecopedia, nil
}
