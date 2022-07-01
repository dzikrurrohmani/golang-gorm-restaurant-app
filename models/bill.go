package model

import (
	"encoding/json"
	"errors"
	"time"

	"gorm.io/gorm"
)

type Bill struct {
	TransDate   time.Time
	CustomerID  uint
	TableID     uint
	TransTypeID uint
	gorm.Model
	Customer    Customer
	Table       Table
	TransType   TransType
}

func (Bill) TableName() string {
	return "t_bill"
}

func (c *Bill) ToString() (string, error) {
	bill, err := json.MarshalIndent(c, "", " ")
	if err != nil {
		return "", errors.New("error in converting to json form")
	}
	return string(bill), nil
}
