package repository

import (
	"errors"
	model "livecode-gorm-wmb/models"

	"gorm.io/gorm"
)

type TableRepository interface {
	Create(table *model.Table) error
	CreateBulk(tables []model.Table) error
	FindById(id uint) (model.Table, error)
	Update(table *model.Table, by map[string]interface{}) error
}

type tableRepository struct {
	db *gorm.DB
}

func (d *tableRepository) Create(table *model.Table) error {
	result := d.db.Create(table)
	return result.Error
}
func (d *tableRepository) CreateBulk(table []model.Table) error {
	result := d.db.Create(table)
	return result.Error
}
func (d *tableRepository) FindById(id uint) (model.Table, error) {
	var table model.Table
	result := d.db.First(&table, id)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return table, nil
		} else {
			return table, err
		}
	}
	return table, nil
}

func (d *tableRepository) Update(table *model.Table, by map[string]interface{}) error {
	result := d.db.Model(table).Updates(by)
	if err := result.Error; err != nil {
		return err
	}
	return nil
}

func NewTableRepository(db *gorm.DB) TableRepository {
	repo := new(tableRepository)
	repo.db = db
	return repo
}
