package model

import (
	"encoding/json"
	"errors"

	"gorm.io/gorm"
)

type BillDetail struct {
	BillID      uint
	MenuPriceID uint
	Qty         int
	gorm.Model
	Bill      Bill
	MenuPrice MenuPrice
}

func (BillDetail) TableName() string {
	return "t_billDetail"
}

func (c *BillDetail) ToString() (string, error) {
	billDetail, err := json.MarshalIndent(c, "", " ")
	if err != nil {
		return "", errors.New("error in converting to json form")
	}
	return string(billDetail), nil
}
