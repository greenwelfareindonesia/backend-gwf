package repository

import (
	"errors"
	"greenwelfare/entity"

	"gorm.io/gorm"
)

type RepositoryEcopedia interface {
	FindAll() ([]*entity.Ecopedia, error)
	FindById(id int) (*entity.Ecopedia, error)
	FindBySlug(slug string) (*entity.Ecopedia, error)
	CreateImage(ecopedia *entity.EcopediaImage) (error)
	Create(ecopedia *entity.Ecopedia) (*entity.Ecopedia, error)
	DeleteEcopedia(ecopedia *entity.Ecopedia)  error
	Update(ecopedia *entity.Ecopedia) (*entity.Ecopedia, error)
	DeleteImages(ecopediaID int) error
	// CreateIsLike(like IsLike) error
}

type repository_ecopedia struct {
	db *gorm.DB
}

func NewRepositoryEcopedia(db *gorm.DB) *repository_ecopedia {
	return &repository_ecopedia{db}
}

// func (r *repository_ecopedia) CreateIsLike(like IsLike) error {
// 	err := r.db.Find(&like).Error
//     if err != nil {
//         return err
//     }
//     return nil
// }



func (r *repository_ecopedia) CreateImage(ecopedia *entity.EcopediaImage) (error) {
	err := r.db.Create(&ecopedia).Error
	return  err
	
}

func (r *repository_ecopedia) DeleteImages(ecopediaID int) error {
    err := r.db.Where("ecopedia_id = ?", ecopediaID).Delete(&entity.EcopediaImage{}).Error
    if err != nil {
        return err
    }

    return nil
}



func (r *repository_ecopedia) FindByUserId (userId int) (*entity.Ecopedia, error) {
	var eco *entity.Ecopedia

	err := r.db.Preload("User").Where("id = ?", userId).Find(eco).Error
	if err != nil {
		return eco, err
	}
	return eco, nil
}

func (r *repository_ecopedia) Update(ecopedia *entity.Ecopedia) (*entity.Ecopedia, error) {
	err := r.db.Save(&ecopedia).Error
	if err != nil {
		return ecopedia, err
	}

	return ecopedia, nil

}

func (r *repository_ecopedia) DeleteEcopedia(ecopedia *entity.Ecopedia) error{
	err := r.db.Delete(&ecopedia).Error
	if err != nil {
		return  err
	}
	return nil
}

func (r *repository_ecopedia) FindAll() ([]*entity.Ecopedia, error) {
	var ecopedias []*entity.Ecopedia
	err := r.db.Preload("FileName").Find(&ecopedias).Error
	if err != nil {
		return ecopedias, err
	}
	return ecopedias, nil
}

// func (r *repository_ecopedia) FindById(id int) (Ecopedia, error) {
// 	var ecopedia Ecopedia
// 	err := r.db.Preload("FileName").Preload("Comment").Preload("Comment.User").Find(&ecopedia, id).Error
// 	if err != nil {
// 		return ecopedia, err
// 	}
// 	return ecopedia, nil
// }

func (r *repository_ecopedia) FindById(id int) (*entity.Ecopedia, error) {
	var ecopedia *entity.Ecopedia
	err := r.db.Preload("FileName").Find(&ecopedia, id).Error
	if err != nil {
		return ecopedia, err
	}
	return ecopedia, nil
}


func (r *repository_ecopedia) FindBySlug(slug string) (*entity.Ecopedia, error) {
	var ecopedia *entity.Ecopedia

	err := r.db.Where("slug = ?", slug).Preload("FileName").Find(&ecopedia).Error

	if err != nil {
		return ecopedia, err
	}
	if ecopedia.Slug == "" {
        return ecopedia, errors.New("slug not found")
    }
	
	return ecopedia, nil

}

func (r *repository_ecopedia) Create(ecopedia *entity.Ecopedia) (*entity.Ecopedia, error) {
	err := r.db.Create(&ecopedia).Error
	if err != nil {
		return ecopedia, err
	}
	return ecopedia, nil
}