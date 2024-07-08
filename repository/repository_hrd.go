package repository

import (
	"greenwelfare/entity"

	"gorm.io/gorm"
)

type RepositoryHRD interface {
	Save(hrd *entity.HRD) (*entity.HRD, error)
	FindByID(ID int) (*entity.HRD, error)
	FindBySlug(slug string) (*entity.HRD, error)
	FindAll() ([]*entity.HRD, error)
	Update(hrd *entity.HRD) (*entity.HRD, error)
	Delete(hrd *entity.HRD) (*entity.HRD, error)
}

type repository_hrd struct {
	db *gorm.DB
}

func NewRepositoryHRD(db *gorm.DB) *repository_hrd {
	return &repository_hrd{db}
}

func (r *repository_hrd) Save(hrd *entity.HRD) (*entity.HRD, error) {
	if err := r.db.Create(&hrd).Error; err != nil {
		return hrd, err
	}

	return hrd, nil
}

func (r *repository_hrd) FindById(ID int) (*entity.HRD, error) {
	var hrd *entity.HRD

	if err := r.db.Where("id = ?", ID).Find(&hrd).Error; err != nil {
		return hrd, err
	}

	return hrd, nil
}

func (r *repository_hrd) FindBySlug(slug string) (*entity.HRD, error) {
	var hrd *entity.HRD
	
	if err := r.db.Where("slug = ?", slug).Find(&hrd).Error; err != nil {
		return hrd, err
	}
	
	return hrd, nil
}

func(r *repository_hrd) FindAll() ([]*entity.HRD, error) {
	var hrds []*entity.HRD

	if err := r.db.Order("id DESC").Find(&hrds).Error; err != nil {
		return nil, err
	}
	
	return hrds, nil
}

func(r *repository_hrd) Update(hrd *entity.HRD) (*entity.HRD, error) {
	if err := r.db.Save(&hrd).Error; err != nil {
		return nil, err
	}

	return hrd, nil
}

func(r *repository_hrd) Delete(hrd *entity.HRD) (*entity.HRD, error) {
	if err := r.db.Delete(&hrd).Error; err != nil {
		return nil, err
	}

	return hrd, nil
}
