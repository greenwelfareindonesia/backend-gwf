package veganguide

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Veganguide, error)
	FindById(id int) (Veganguide, error)
	Create(veganguide Veganguide) (Veganguide, error)
	DeleteVeganguide(veganguide Veganguide) (Veganguide, error)
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
	err := r.db.Find(&veganguide, id).Error
	if err != nil {
		return veganguide, err
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
