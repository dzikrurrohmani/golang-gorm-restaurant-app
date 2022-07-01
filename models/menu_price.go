package model

import (
	"encoding/json"
	"errors"

	"gorm.io/gorm"
)

type MenuPrice struct {
	Price float32 `gorm:"not null"`
	MenuID uint
	Menu Menu
	gorm.Model

}

func (MenuPrice) TableName() string {
	return "m_menu_price"
}

func (c *MenuPrice) ToString() (string, error) {
	menuPrice, err := json.MarshalIndent(c, "", " ")
	if err != nil {
		return "", errors.New("error in converting to json form")
	}
	return string(menuPrice), nil
}
