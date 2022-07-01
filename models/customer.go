package model

import (
	"encoding/json"
	"errors"

	"gorm.io/gorm"
)

type Customer struct {
	CustomerName  string      `gorm:"size:50;not null"`
	MobilePhoneNo string      `gorm:"unique; size:13"`
	IsMember      bool        `gorm:"default:false"`
	Discounts     []*Discount `gorm:"many2many:m_customer_discount"`
	gorm.Model
}

func (Customer) TableName() string {
	return "m_customer"
}

func (c *Customer) ToString() (string, error) {
	customer, err := json.MarshalIndent(c, "", " ")
	if err != nil {
		return "", errors.New("error in converting to json form")
	}
	return string(customer), nil
}
