package repository

import (
	"greenwelfare/entity"

	"gorm.io/gorm"
)

type RepositoryUser interface {
	//create User
	Save(user *entity.User) (*entity.User, error)
	FindById(ID int) (*entity.User, error)
	FindBySlug(slug string) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
	Update(user *entity.User) (*entity.User, error)
	Delete(user *entity.User) (*entity.User, error)
}

type repository_user struct {
	db *gorm.DB
}

func NewRepositoryUser(db *gorm.DB) *repository_user {
	return &repository_user{db}
}

func (r *repository_user) Save(user *entity.User) (*entity.User, error) {
	err := r.db.Create(&user).Error

	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository_user) FindBySlug(slug string) (*entity.User, error) {
	var user *entity.User
	err := r.db.Where("slug = ?", slug).Find(&user).Error

	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository_user) FindByEmail(email string) (*entity.User, error) {
	var user *entity.User
	err := r.db.Where("email = ?", email).Find(&user).Error

	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository_user) FindById(ID int) (*entity.User, error) {
	var user *entity.User

	err := r.db.Where("id = ?", ID).Find(&user).Error

	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository_user) Update(user *entity.User) (*entity.User, error) {
	err := r.db.Save(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil

}

func (r *repository_user) Delete(user *entity.User) (*entity.User, error) {
	err := r.db.Delete(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
