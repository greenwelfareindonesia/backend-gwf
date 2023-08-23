package artikel

import "gorm.io/gorm"

type Repository interface {
	//create User
	Save(karya Artikel) (Artikel, error)
	FindById(ID int) (Artikel, error)
	FindAll() ([]Artikel, error)
	FindByKarya(Karya int) ([]Artikel, error)
	FindByTags(tags int) ([]Artikel, error)
	Update(artikel Artikel) (Artikel, error)
	Delete(karya Artikel) (Artikel, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Artikel, error) {
	var karya []Artikel

	err := r.db.Order("id DESC").Find(&karya).Error
	if err != nil {
		return karya, err
	}
	return karya, nil
}

func (r *repository) Save(karya Artikel) (Artikel, error) {
	err := r.db.Create(&karya).Error

	if err != nil {
		return karya, err
	}
	return karya, nil
}

func (r *repository) FindByTags(tags int) ([]Artikel, error) {
	var karya []Artikel

	err := r.db.Where("tags_id = ?", tags).Find(&karya).Error

	if err != nil {
		return karya, err
	}
	return karya, nil
}

func (r *repository) FindByKarya(Karya int) ([]Artikel, error) {
	var karya []Artikel

	err := r.db.Where("karya_news_id = ?", Karya).Find(&karya).Error
	if err != nil {
		return karya, err
	}
	return karya, nil
}

func (r *repository) FindById(ID int) (Artikel, error) {
	var karya Artikel

	err := r.db.Where("id = ?", ID).Find(&karya).Error

	if err != nil {
		return karya, err
	}
	return karya, nil
}

func (r *repository) Update(karya Artikel) (Artikel, error) {
	err := r.db.Save(&karya).Error
	if err != nil {
		return karya, err
	}

	return karya, nil

}

func (r *repository) Delete(karya Artikel) (Artikel, error) {
	err := r.db.Delete(&karya).Error
	if err != nil {
		return karya, err
	}

	return karya, nil
}