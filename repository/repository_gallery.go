package repository

import (
	"errors"
	"fmt"
	"greenwelfare/entity"

	"gorm.io/gorm"
)

type RepositoryGallery interface {
	Create(gallery *entity.Gallery) (*entity.Gallery, error)
	FindAll() ([]*entity.Gallery, error)
	FindById(ID int) (*entity.Gallery, error)
	FindBySlug(slug string) (*entity.Gallery, error)
	CreateImage(gallery *entity.GalleryImages) error
	DeleteImages(galleryID int) error
	Update(gallery *entity.Gallery) (*entity.Gallery, error)
	Delete(gallery *entity.Gallery) (*entity.Gallery, error)
}

type repository_galery struct {
	db *gorm.DB
}

func NewRepositoryGallery(db *gorm.DB) *repository_galery {
	return &repository_galery{db}
}

func (r *repository_galery) DeleteImages(galleryID int) error {
	err := r.db.Where("gallery_id = ?", galleryID).Delete(&entity.GalleryImages{}).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repository_galery) FindBySlug(slug string) (*entity.Gallery, error) {
	var gallery *entity.Gallery

	err := r.db.Where("slug = ?", slug).Preload("FileName").Find(&gallery).Error

	if err != nil {
		return gallery, err
	}
	if gallery.Slug == "" {
		return gallery, errors.New("slug not found")
	}

	return gallery, nil

}

func (r *repository_galery) Create(gallery *entity.Gallery) (*entity.Gallery, error) {
	err := r.db.Create(&gallery).Error

	if err != nil {
		return gallery, err
	}
	return gallery, nil
}

func (r *repository_galery) CreateImage(gallery *entity.GalleryImages) error {
	err := r.db.Create(&gallery).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository_galery) FindAll() ([]*entity.Gallery, error) {
	var gallery []*entity.Gallery

	err := r.db.Preload("FileName").Find(&gallery).Error

	if err != nil {
		return gallery, err
	}
	return gallery, nil
}

func (r *repository_galery) FindById(ID int) (*entity.Gallery, error) {
	var gallery *entity.Gallery

	err := r.db.Where("id = ?", ID).Find(&gallery).Error

	if err != nil {
		return gallery, err
	}

	if gallery.ID == 0 {
		return gallery, errors.New(fmt.Sprintf("gallery with id %d not found", ID))
	}
	return gallery, nil
}

func (r *repository_galery) Update(gallery *entity.Gallery) (*entity.Gallery, error) {
	err := r.db.Save(&gallery).Error

	if err != nil {
		return gallery, err
	}
	return gallery, nil
}

func (r *repository_galery) Delete(gallery *entity.Gallery) (*entity.Gallery, error) {
	err := r.db.Delete(&gallery).Error

	if err != nil {
		return gallery, err
	}
	return gallery, nil
}
