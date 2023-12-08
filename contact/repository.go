package contact

import (
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	Submit(contact_submission Contact) (Contact, error)
	FindAll() ([]Contact, error)
	FindById(ID int) (Contact, error)
	Delete(contact_submission Contact) (Contact, error)
	FindBySlug(slug string) (Contact, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindBySlug(slug string) (Contact, error) {
	var contact Contact

	err := r.db.Where("slug = ?", slug).Find(&slug).Error

	if err != nil {
		return contact, err
	}
	if contact.Slug == "" {
        return contact, errors.New("slug not found")
    }
	
	return contact, nil
}

func (r *repository) Submit(contact_submission Contact) (Contact, error) {
	err := r.db.Create(&contact_submission).Error

	if err != nil {
		return contact_submission, err
	}
	return contact_submission, nil
}

func (r *repository) FindAll() ([]Contact, error) {
	var contact_submission []Contact

	err := r.db.Find(&contact_submission).Error

	if err != nil {
		return contact_submission, err
	}
	return contact_submission, nil
}

func (r *repository) FindById(ID int) (Contact, error) {
	var contact_submission Contact

	err := r.db.Where("id = ?", ID).Find(&contact_submission).Error

	if err != nil {
		return contact_submission, err
	}
	return contact_submission, nil
}

func (r *repository) FIndByName(Name string) (Contact, error) {
	var contact_submission Contact

	err := r.db.Where("name = ?", Name).Find(&contact_submission).Error

	if err != nil {
		return contact_submission, err
	}
	return contact_submission, nil
}

func (r *repository) FindByEmail(Email string) (Contact, error) {
	var contact_submission Contact

	err := r.db.Where("email = ?", Email).Find(&contact_submission).Error

	if err != nil {
		return contact_submission, err
	}
	return contact_submission, nil
}

func (r *repository) Delete(contact_submission Contact) (Contact, error) {
	err := r.db.Delete(&contact_submission).Error

	if err != nil {
		return contact_submission, err
	}
	return contact_submission, nil
}
