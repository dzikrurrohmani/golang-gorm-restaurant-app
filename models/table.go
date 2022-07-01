package model

import (
	"encoding/json"
	"errors"

	"gorm.io/gorm"
)

type Table struct {
	TableDescription string
	IsAvailable bool `gorm:"default:true"`
	gorm.Model
}

func (Table) TableName() string {
	return "m_table"
}

func (c *Table) ToString() (string, error) {
	table, err := json.MarshalIndent(c, "", " ")
	if err != nil {
		return "", errors.New("error in converting to json form")
	}
	return string(table), nil
}
