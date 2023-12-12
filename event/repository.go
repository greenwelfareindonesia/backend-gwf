package event

import (
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	//create User
	Save(karya Event) (Event, error)
	FindById(ID int) (Event, error)
	FindAll() ([]Event, error)
	FindBySlug(slug string) (Event, error)
	Update(artikel Event) (Event, error)
	Delete(karya Event) (Event, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindBySlug(slug string) (Event, error) {
	var event Event

	err := r.db.Where("slug = ?", slug).Find(&event).Error

	if err != nil {
		return event, err
	}
	if event.Slug == "" {
        return event, errors.New("slug not found")
    }
	
	return event, nil

}

func (r *repository) FindAll() ([]Event, error) {
	var karya []Event

	err := r.db.Order("id DESC").Find(&karya).Error
	if err != nil {
		return karya, err
	}
	return karya, nil
}

func (r *repository) Save(karya Event) (Event, error) {
	err := r.db.Create(&karya).Error

	if err != nil {
		return karya, err
	}
	return karya, nil
}

func (r *repository) FindByTags(tags int) ([]Event, error) {
	var karya []Event

	err := r.db.Where("tags_id = ?", tags).Find(&karya).Error

	if err != nil {
		return karya, err
	}
	return karya, nil
}

func (r *repository) FindByKarya(Karya int) ([]Event, error) {
	var karya []Event

	err := r.db.Where("karya_news_id = ?", Karya).Find(&karya).Error
	if err != nil {
		return karya, err
	}
	return karya, nil
}

func (r *repository) FindById(ID int) (Event, error) {
	var karya Event

	err := r.db.Where("id = ?", ID).Find(&karya).Error

	if err != nil {
		return karya, err
	}
	return karya, nil
}

func (r *repository) Update(karya Event) (Event, error) {
	err := r.db.Save(&karya).Error
	if err != nil {
		return karya, err
	}
	return karya, nil
}

func (r *repository) Delete(karya Event) (Event, error) {
	err := r.db.Delete(&karya).Error
	if err != nil {
		return karya, err
	}

	return karya, nil
}
