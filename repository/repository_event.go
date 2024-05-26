package repository

import (
	"errors"
	"greenwelfare/entity"

	"gorm.io/gorm"
)

type RepositoryEvent interface {
	//create User
	Save(karya *entity.Event) (*entity.Event, error)
	FindById(ID int) (*entity.Event, error)
	FindAll() ([]*entity.Event, error)
	FindBySlug(slug string) (*entity.Event, error)
	Update(artikel *entity.Event) (*entity.Event, error)
	Delete(karya *entity.Event) (*entity.Event, error)
}

type repository_event struct {
	db *gorm.DB
}

func NewRepositoryEvent(db *gorm.DB) *repository_event {
	return &repository_event{db}
}

func (r *repository_event) FindBySlug(slug string) (*entity.Event, error) {
	var event *entity.Event

	err := r.db.Where("slug = ?", slug).Find(&event).Error

	if err != nil {
		return event, err
	}
	if event.Slug == "" {
        return event, errors.New("slug not found")
    }
	
	return event, nil

}

func (r *repository_event) FindAll() ([]*entity.Event, error) {
	var karya []*entity.Event

	err := r.db.Order("id DESC").Find(&karya).Error
	if err != nil {
		return karya, err
	}
	return karya, nil
}

func (r *repository_event) Save(karya *entity.Event) (*entity.Event, error) {
	err := r.db.Create(&karya).Error

	if err != nil {
		return karya, err
	}
	return karya, nil
}

func (r *repository_event) FindByTags(tags int) ([]*entity.Event, error) {
	var karya []*entity.Event

	err := r.db.Where("tags_id = ?", tags).Find(&karya).Error

	if err != nil {
		return karya, err
	}
	return karya, nil
}

func (r *repository_event) FindByKarya(Karya int) ([]*entity.Event, error) {
	var karya []*entity.Event

	err := r.db.Where("karya_news_id = ?", Karya).Find(&karya).Error
	if err != nil {
		return karya, err
	}
	return karya, nil
}

func (r *repository_event) FindById(ID int) (*entity.Event, error) {
	var karya *entity.Event

	err := r.db.Where("id = ?", ID).Find(&karya).Error

	if err != nil {
		return karya, err
	}
	return karya, nil
}

func (r *repository_event) Update(karya *entity.Event) (*entity.Event, error) {
	err := r.db.Save(&karya).Error
	if err != nil {
		return karya, err
	}
	return karya, nil
}

func (r *repository_event) Delete(karya *entity.Event) (*entity.Event, error) {
	err := r.db.Delete(&karya).Error
	if err != nil {
		return karya, err
	}

	return karya, nil
}
