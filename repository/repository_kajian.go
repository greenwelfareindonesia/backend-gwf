package repository

import (
	"fmt"
	"greenwelfare/entity"
	"math/rand"
	"strings"
	"time"

	"gorm.io/gorm"
)

type RepositoryKajian interface {
	Save(kajian *entity.Kajian) (*entity.Kajian, error)
	SaveImage(kajianImage *entity.KajianImage) error
	FindByID(id int) (*entity.Kajian, error)
	FindBySlug(slug string) (*entity.Kajian, error)
	FindAll() ([]*entity.Kajian, error)
	Update(kajian *entity.Kajian) (*entity.Kajian, error)
	Delete(kajian *entity.Kajian) error
	DeleteImages(kajian *entity.Kajian) error
}

type repository_kajian struct {
	db *gorm.DB
}

func NewRepositoryKajian(db *gorm.DB) *repository_kajian {
	return &repository_kajian{db}
}

func (r *repository_kajian) Save(kajian *entity.Kajian) (*entity.Kajian, error) {
	slugName := strings.ReplaceAll(strings.ToLower(kajian.Title), " ", "-")
	randomNumber := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(1000000)
	kajian.Slug = fmt.Sprintf("%s-%d", slugName, randomNumber)

	if err := r.db.Create(&kajian).Error; err != nil {
		return nil, err
	}

	return kajian, nil
}

func (r *repository_kajian) SaveImage(kajianImage *entity.KajianImage) error {
	if err := r.db.Create(&kajianImage).Error; err != nil {
		return err
	}

	return nil
}

func (r *repository_kajian) FindByID(id int) (*entity.Kajian, error) {
	var kajian *entity.Kajian

	if err := r.db.Where("id = ?", id).Preload("Images").Find(&kajian).Error; err != nil {
		return nil, err
	}

	return kajian, nil
}

func (r *repository_kajian) FindBySlug(slug string) (*entity.Kajian, error) {
	var kajian *entity.Kajian

	if err := r.db.Where("slug = ?", slug).Preload("Images").Find(&kajian).Error; err != nil {
		return nil, err
	}

	return kajian, nil
}

func (r *repository_kajian) FindAll() ([]*entity.Kajian, error) {
	var kajians []*entity.Kajian

	if err := r.db.Preload("Images").Find(&kajians).Error; err != nil {
		return nil, err
	}

	return kajians, nil
}

func (r *repository_kajian) Update(kajian *entity.Kajian) (*entity.Kajian, error) {
	if err := r.db.Save(&kajian).Error; err != nil {
		return nil, err
	}

	return kajian, nil
}

func (r *repository_kajian) Delete(kajian *entity.Kajian) error {
	if err := r.db.Delete(&kajian).Error; err != nil {
		return err
	}

	return nil
}

func (r *repository_kajian) DeleteImages(kajian *entity.Kajian) error {
	if err := r.db.Where("kajian_id = ?", kajian.ID).Delete(&entity.KajianImage{}).Error; err != nil {
		return err
	}

	return nil
}
