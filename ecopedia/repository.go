package ecopedia

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Ecopedia, error)
	FindById(id int) (Ecopedia, error)
	Create(ecopedia Ecopedia) (Ecopedia, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Ecopedia, error) {
	var ecopedias []Ecopedia
	err := r.db.Find(&ecopedias).Error
	if err != nil {
		return ecopedias, err
	}
	return ecopedias, nil
}

func (r *repository) FindById(id int) (Ecopedia, error) {
	var ecopedia Ecopedia
	err := r.db.Find(&ecopedia, id).Error
	if err != nil {
		return ecopedia, err
	}
	return ecopedia, nil
}

func (r *repository) Create(ecopedia Ecopedia) (Ecopedia, error) {
	err := r.db.Create(&ecopedia).Error
	if err != nil {
		return ecopedia, err
	}
	return ecopedia, nil
}

// func (r *repository) Create(ecopedia Ecopedia) (Ecopedia, error) {
// 	err := r.db.Create(&ecopedia).Error
// if err != nil {
// 	return ecopedia, err
// }
// 	return ecopedia, nil
// }

// func (r *repository) FindAll() ([]Ecopedia, error) {
// 	var ecopedias []Ecopedia
// 	err := r.db.Find(&ecopedias).Error
// 	if err != nil {
// 		return ecopedias, err
// 	}
// 	return ecopedias, nil
// }

// func (r *repository) FindById(id int) (Ecopedia, error) {
// 	var ecopedia Ecopedia
// 	err := r.db.Find(&ecopedia, id).Error
// 	return ecopedia, err
// }
