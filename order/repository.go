package transactions

import (
	"greenwelfare/entity"
	"greenwelfare/product"

	"gorm.io/gorm"
)

type RepositoryOrder interface {
	FindAll() ([]Order, error)
	Save(order Order) (Order, error)
	FindById(ID int) (Order, error)
	Update(order Order) (Order, error)
	Delete(order Order) (Order, error)
	FindByUserId(productID int, userID int) ([]Order, error)
}

type repositoryOrder struct {
	db *gorm.DB
}

func NewRepositoryOrder(db *gorm.DB) *repositoryOrder {
	return &repositoryOrder{db}
}

func (r *repositoryOrder) FindAll() ([]Order, error) {
	var order []Order

	err := r.db.Preload("Product").Preload("User").Find(&order).Error

	if err != nil {
		return order, err
	}

	return order, nil
}

func (r *repositoryOrder) Save(order Order) (Order, error) {
	err := r.db.Create(&order).Error

	if err != nil {
		return order, err
	}
	return order, nil
}

func (r *repositoryOrder) FindByUserId(productID int, userID int) ([]Order, error) {
	var order []Order

	err := r.db.Joins("Product", r.db.Where(&product.Products{ID: productID})).Joins("User", r.db.Where(&entity.User{ID: userID})).Find(&order).Error

	if err != nil {
		return order, err
	}
	return order, nil
}

func (r *repositoryOrder) FindById(ID int) (Order, error) {
	var order Order

	err := r.db.Where("id = ?", ID).Find(&order).Error

	if err != nil {
		return order, err
	}
	return order, nil
}

func (r *repositoryOrder) Update(order Order) (Order, error) {
	err := r.db.Save(&order).Error
	if err != nil {
		return order, err
	}

	return order, nil

}

func (r *repositoryOrder) Delete(order Order) (Order, error) {
	err := r.db.Delete(&order).Error
	if err != nil {
		return order, err
	}

	return order, nil
}
