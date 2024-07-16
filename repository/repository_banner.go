package repository

import (
	"context"
	"greenwelfare/entity"

	"gorm.io/gorm"
)

type RepositoryBanner interface {
	CreateBanner(ctx context.Context, banner entity.Banner) (entity.Banner, error)
}

type repository_banner struct {
	db *gorm.DB
}

func NewRepositoryBannder(db *gorm.DB) *repository_banner {
	return &repository_banner{db}
}

func (r *repository_banner) CreateBanner(ctx context.Context, banner entity.Banner) (entity.Banner, error) {
	if err := r.db.WithContext(ctx).Table("banners").Save(&banner).Error; err != nil {
		return entity.Banner{}, err
	}

	return banner, nil
}
