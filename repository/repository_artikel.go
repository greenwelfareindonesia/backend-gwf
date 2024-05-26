package repository

import (
	"errors"
	"greenwelfare/entity"

	"gorm.io/gorm"
)

type RepositoryArtikel interface {
	//create User
	Save(karya *entity.Artikel) (*entity.Artikel, error)
	FindById(ID int) (*entity.Artikel, error)
	FindBySlug(slug string) (*entity.Artikel, error)
	FindAll() ([]*entity.Artikel, error)
	Update(artikel *entity.Artikel) (*entity.Artikel, error)
	Delete(karya *entity.Artikel) (*entity.Artikel, error)
}

type repository_artikel struct {
	db *gorm.DB
}

func NewRepositoryArtikel(db *gorm.DB) *repository_artikel {
	return &repository_artikel{db}
}

func (r *repository_artikel) FindAll() ([]*entity.Artikel, error) {
	var karya []*entity.Artikel

	err := r.db.Order("id DESC").Find(&karya).Error
	if err != nil {
		return karya, err
	}
	return karya, nil
}

func (r *repository_artikel) Save(karya *entity.Artikel) (*entity.Artikel, error) {
	err := r.db.Create(&karya).Error

	if err != nil {
		return karya, err
	}
	return karya, nil
}

func (r *repository_artikel) FindById(ID int) (*entity.Artikel, error) {
	var karya *entity.Artikel

	err := r.db.Where("id = ?", ID).Find(&karya).Error

	if err != nil {
		return karya, err
	}
	return karya, nil
}

func (r *repository_artikel) FindBySlug(slug string) (*entity.Artikel, error) {
	var artikel *entity.Artikel

	err := r.db.Where("slug = ?", slug).Find(&artikel).Error

	if err != nil {
		return artikel, err
	}
	if artikel.Slug == "" {
        return artikel, errors.New("slug not found")
    }
	
	return artikel, nil

}

func (r *repository_artikel) Update(karya *entity.Artikel) (*entity.Artikel, error) {
	err := r.db.Save(&karya).Error
	if err != nil {
		return karya, err
	}

	return karya, nil

}

func (r *repository_artikel) Delete(karya *entity.Artikel) (*entity.Artikel, error) {
	err := r.db.Delete(&karya).Error
	if err != nil {
		return karya, err
	}

	return karya, nil
}