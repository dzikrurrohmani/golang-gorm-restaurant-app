package model

import (
	"encoding/json"
	"errors"
	"time"
)

type Bill struct {
	BaseModel   BaseModel `gorm:"embedded"`
	TransDate   time.Time
	CustomerID  uint
	TableID      uint
	TransTypeID string
	TransType   TransType
	BillDetails []BillDetail
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
