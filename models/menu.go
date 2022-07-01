package model

import (
	"encoding/json"
	"errors"

	"gorm.io/gorm"
)

type Menu struct {
	MenuName string `gorm:"size:50;not null"`
	gorm.Model
}

func (Menu) TableName() string {
	return "m_menu"
}

func (c *Menu) ToString() (string, error) {
	menu, err := json.MarshalIndent(c, "", " ")
	if err != nil {
		return "", errors.New("error in converting to json form")
	}
	return string(menu), nil
}
