package ecopedia

import (
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]Ecopedia, error)
	FindById(id int) (Ecopedia, error)
	FindBySlug(slug string) (Ecopedia, error)
	CreateImage(ecopedia EcopediaImage) (error)
	Create(ecopedia Ecopedia) (Ecopedia, error)
	DeleteEcopedia(ecopedia Ecopedia)  error
	Update(ecopedia Ecopedia) (Ecopedia, error)
	DeleteImages(ecopediaID int) error
	// CreateIsLike(like IsLike) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// func (r *repository) CreateIsLike(like IsLike) error {
// 	err := r.db.Find(&like).Error
//     if err != nil {
//         return err
//     }
//     return nil
// }



func (r *repository) CreateImage(ecopedia EcopediaImage) (error) {
	err := r.db.Create(&ecopedia).Error
	return  err
	
}

func (r *repository) DeleteImages(ecopediaID int) error {
    err := r.db.Where("ecopedia_id = ?", ecopediaID).Delete(&EcopediaImage{}).Error
    if err != nil {
        return err
    }

    return nil
}



func (r *repository) FindByUserId (userId int) (Ecopedia, error) {
	var eco Ecopedia

	err := r.db.Preload("User").Where("id = ?", userId).Find(eco).Error
	if err != nil {
		return eco, err
	}
	return eco, nil
}

func (r *repository) Update(ecopedia Ecopedia) (Ecopedia, error) {
	err := r.db.Save(&ecopedia).Error
	if err != nil {
		return ecopedia, err
	}

	return ecopedia, nil

}

func (r *repository) DeleteEcopedia(ecopedia Ecopedia) error{
	err := r.db.Delete(&ecopedia).Error
	if err != nil {
		return  err
	}
	return nil
}

func (r *repository) FindAll() ([]Ecopedia, error) {
	var ecopedias []Ecopedia
	err := r.db.Preload("FileName").Find(&ecopedias).Error
	if err != nil {
		return ecopedias, err
	}
	return ecopedias, nil
}

// func (r *repository) FindById(id int) (Ecopedia, error) {
// 	var ecopedia Ecopedia
// 	err := r.db.Preload("FileName").Preload("Comment").Preload("Comment.User").Find(&ecopedia, id).Error
// 	if err != nil {
// 		return ecopedia, err
// 	}
// 	return ecopedia, nil
// }

func (r *repository) FindById(id int) (Ecopedia, error) {
	var ecopedia Ecopedia
	err := r.db.Preload("FileName").Find(&ecopedia, id).Error
	if err != nil {
		return ecopedia, err
	}
	return ecopedia, nil
}


func (r *repository) FindBySlug(slug string) (Ecopedia, error) {
	var ecopedia Ecopedia

	err := r.db.Where("slug = ?", slug).Preload("FileName").Find(&ecopedia).Error

	if err != nil {
		return ecopedia, err
	}
	if ecopedia.Slug == "" {
        return ecopedia, errors.New("slug not found")
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