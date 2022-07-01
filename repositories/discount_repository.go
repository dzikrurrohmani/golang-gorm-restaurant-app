package repository

import (
	"errors"
	model "livecode-gorm-wmb/models"

	"gorm.io/gorm"
)

type DiscountRepository interface {
	Create(disc *model.Discount) error
	FindById(id uint) (model.Discount, error)
}

type discountRepository struct {
	db *gorm.DB
}

func (d *discountRepository) Create(disc *model.Discount) error {
	result := d.db.Create(disc)
	return result.Error
}
func (d *discountRepository) FindById(id uint) (model.Discount, error) {
	var disc model.Discount
	result := d.db.First(&disc, id)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return disc, nil
		} else {
			return disc, err
		}
	}
	return disc, nil
}

func NewDiscountRepository(db *gorm.DB) DiscountRepository {
	repo := new(discountRepository)
	repo.db = db
	return repo
}
