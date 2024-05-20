package repository

import (
	"errors"
	"greenwelfare/entity"

	"gorm.io/gorm"
)

type RepositoryContact interface {
	Submit(contact_submission *entity.Contact) (*entity.Contact, error)
	FindAll() ([]*entity.Contact, error)
	FindById(ID int) (*entity.Contact, error)
	Delete(contact_submission *entity.Contact) (*entity.Contact, error)
	FindBySlug(slug string) (*entity.Contact, error)
}

type repository_contact struct {
	db *gorm.DB
}

func NewRepositoryContact(db *gorm.DB) *repository_contact {
	return &repository_contact{db}
}

func (r *repository_contact) FindBySlug(slug string) (*entity.Contact, error) {
	var contact *entity.Contact

	err := r.db.Where("slug = ?", slug).Find(&contact).Error

	if err != nil {
		return contact, err
	}
	if contact.Slug == "" {
		return contact, errors.New("slug not found")
	}

	return contact, nil
}

func (r *repository_contact) Submit(contact_submission *entity.Contact) (*entity.Contact, error) {
	err := r.db.Create(&contact_submission).Error

	if err != nil {
		return contact_submission, err
	}
	return contact_submission, nil
}

func (r *repository_contact) FindAll() ([]*entity.Contact, error) {
	var contact_submission []*entity.Contact

	err := r.db.Find(&contact_submission).Error

	if err != nil {
		return contact_submission, err
	}
	return contact_submission, nil
}

func (r *repository_contact) FindById(ID int) (*entity.Contact, error) {
	var contact_submission *entity.Contact

	err := r.db.Where("id = ?", ID).Find(&contact_submission).Error

	if err != nil {
		return contact_submission, err
	}
	return contact_submission, nil
}

func (r *repository_contact) FIndByName(Name string) (*entity.Contact, error) {
	var contact_submission *entity.Contact

	err := r.db.Where("name = ?", Name).Find(&contact_submission).Error

	if err != nil {
		return contact_submission, err
	}
	return contact_submission, nil
}

func (r *repository_contact) FindByEmail(Email string) (*entity.Contact, error) {
	var contact_submission *entity.Contact

	err := r.db.Where("email = ?", Email).Find(&contact_submission).Error

	if err != nil {
		return contact_submission, err
	}
	return contact_submission, nil
}

func (r *repository_contact) Delete(contact_submission *entity.Contact) (*entity.Contact, error) {
	err := r.db.Delete(&contact_submission).Error

	if err != nil {
		return contact_submission, err
	}
	return contact_submission, nil
}
