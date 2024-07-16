package repository

import (
	"greenwelfare/entity"

	"gorm.io/gorm"
)

type RepositoryHrd interface {
	Save(hrd *entity.Hrd) (*entity.Hrd, error)
	FindByID(ID int) (*entity.Hrd, error)
	FindBySlug(slug string) (*entity.Hrd, error)
	FindAll(hrd *entity.Hrd) ([]*entity.Hrd, error)
	Update(hrd *entity.Hrd) (*entity.Hrd, error)
	Delete(hrd *entity.Hrd) (*entity.Hrd, error)
}

type repository_hrd struct {
	db *gorm.DB
}

func NewRepositoryHRD(db *gorm.DB) *repository_hrd {
	return &repository_hrd{db}
}

func (r *repository_hrd) Save(hrd *entity.Hrd) (*entity.Hrd, error) {
	if err := r.db.Create(&hrd).Error; err != nil {
		return hrd, err
	}

	return hrd, nil
}

func (r *repository_hrd) FindByID(ID int) (*entity.Hrd, error) {
	var hrd *entity.Hrd

	if err := r.db.Where("id = ?", ID).Find(&hrd).Error; err != nil {
		return hrd, err
	}

	return hrd, nil
}

func (r *repository_hrd) FindBySlug(slug string) (*entity.Hrd, error) {
	var hrd *entity.Hrd
	
	if err := r.db.Where("slug = ?", slug).Find(&hrd).Error; err != nil {
		return hrd, err
	}
	
	return hrd, nil
}

func(r *repository_hrd) FindAll(hrd *entity.Hrd) ([]*entity.Hrd, error) {
	var hrds []*entity.Hrd

	if err := r.db.Order("id DESC").Where(hrd).Find(&hrds).Error; err != nil {
		return nil, err
	}
	
	return hrds, nil
}

func(r *repository_hrd) Update(hrd *entity.Hrd) (*entity.Hrd, error) {
	if err := r.db.Save(&hrd).Error; err != nil {
		return nil, err
	}

	return hrd, nil
}

func(r *repository_hrd) Delete(hrd *entity.Hrd) (*entity.Hrd, error) {
	if err := r.db.Delete(&hrd).Error; err != nil {
		return nil, err
	}

	return hrd, nil
}
