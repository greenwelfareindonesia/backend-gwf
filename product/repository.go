package product

import "gorm.io/gorm"

type RepositoryProduct interface {
	FindAll() ([]Products, error)
	Save(product Products) (Products, error)
	FindById(ID int) (Products, error)
	Update(product Products) (Products, error)
	Delete(product Products) (Products, error)
	FindAllProductByCategory(ID int) ([]Products, error)
}

type repositoryProduct struct {
	db *gorm.DB
}

func NewRepositoryProduct(db *gorm.DB) *repositoryProduct {
	return &repositoryProduct{db}
}

func (r *repositoryProduct) FindAllProductByCategory(ID int) ([]Products, error) {
	var product []Products

	err := r.db.Preload("Category").Where("category_id = ? ", ID).Find(&product).Error

	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *repositoryProduct) FindAll() ([]Products, error) {
	var product []Products

	err := r.db.Preload("Category").Find(&product).Error

	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *repositoryProduct) Save(product Products) (Products, error) {
	err := r.db.Create(&product).Error

	if err != nil {
		return product, err
	}
	return product, nil
}

func (r *repositoryProduct) FindById(ID int) (Products, error) {
	var product Products

	err := r.db.Where("id = ?", ID).Find(&product).Error

	if err != nil {
		return product, err
	}
	return product, nil
}

func (r *repositoryProduct) Update(product Products) (Products, error) {
	err := r.db.Save(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil

}

func (r *repositoryProduct) Delete(product Products) (Products, error) {
	err := r.db.Delete(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}
