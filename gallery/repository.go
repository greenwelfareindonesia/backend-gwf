package gallery

import (
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	Create(gallery Gallery) (Gallery, error)
	FindAll() ([]Gallery, error)
	FindById(ID int) (Gallery, error)
	FindBySlug(slug string) (Gallery, error)
	CreateImage(gallery GalleryImages) (error)
	DeleteImages(galleryID int) error
	Update(gallery Gallery) (Gallery, error)
	Delete(gallery Gallery) (Gallery, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) DeleteImages(galleryID int) error {
    err := r.db.Where("gallery_id = ?", galleryID).Delete(&GalleryImages{}).Error
    if err != nil {
        return err
    }

    return nil
}


func (r *repository) FindBySlug(slug string) (Gallery, error) {
	var gallery Gallery

	err := r.db.Where("slug = ?", slug).Preload("FileName").Find(&gallery).Error

	if err != nil {
		return gallery, err
	}
	if gallery.Slug == "" {
        return gallery, errors.New("slug not found")
    }
	
	return gallery, nil

}

func (r *repository) Create(gallery Gallery) (Gallery, error) {
	err := r.db.Create(&gallery).Error

	if err != nil {
		return gallery, err
	}
	return gallery, nil
}

func (r *repository) CreateImage(gallery GalleryImages) (error) {
	err := r.db.Create(&gallery).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) FindAll() ([]Gallery, error) {
	var gallery []Gallery

	err := r.db.Preload("FileName").Find(&gallery).Error

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
