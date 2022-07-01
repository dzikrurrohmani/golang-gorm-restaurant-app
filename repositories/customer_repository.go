package repository

import (
	"errors"
	model "livecode-gorm-wmb/models"

	"gorm.io/gorm"
)

type CustomerRepository interface {
	Create(customer *model.Customer) error
	FindById(id uint) (model.Customer, error)
	FindAllBy(preload string, condition string, searchValue ...interface{}) ([]model.Customer, error)
	Update(customer *model.Customer, by map[string]interface{}) error
	Count() (int64, error)
}
type customerRepository struct {
	db *gorm.DB
}

func (c *customerRepository) Create(customer *model.Customer) error {
	result := c.db.Create(customer)
	return result.Error
}

func (c *customerRepository) FindById(id uint) (model.Customer, error) {
	var customer model.Customer
	result := c.db.First(&customer, id)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return customer, nil
		} else {
			return customer, err
		}
	}
	return customer, nil
}
func (c *customerRepository) FindAllBy(preload string, condition string, searchValue ...interface{}) ([]model.Customer, error) {
	var customers []model.Customer
	if preload == "" {
		result := c.db.Where(condition, searchValue...).Find(&customers)
		if err := result.Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, nil
			} else {
				return nil, err
			}
		}
	} else {
		result := c.db.Preload(preload).Where(condition, searchValue...).Find(&customers)
		if err := result.Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, nil
			} else {
				return nil, err
			}
		}
	}

	return customers, nil
}

func (c *customerRepository) Update(customer *model.Customer, by map[string]interface{}) error {
	result := c.db.Model(customer).Updates(by)
	if err := result.Error; err != nil {
		return err
	}
	return nil
}
func (c *customerRepository) Count() (int64, error) {
	var total int64
	result := c.db.Model(&model.Customer{}).Count(&total)
	if err := result.Error; err != nil {
		return 0, err
	}
	return total, nil
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	repo := new(customerRepository)
	repo.db = db
	return repo
}
