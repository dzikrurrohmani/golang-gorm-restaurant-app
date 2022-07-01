package migration

import (
	model "livecode-gorm-wmb/models"
	util "livecode-gorm-wmb/utils"

	"gorm.io/gorm"
)

func MigrateDb(db *gorm.DB) {
	err := db.AutoMigrate(
		&model.Customer{},
		&model.Menu{},
		&model.Table{},
		&model.Discount{},
		&model.MenuPrice{},
		&model.TransType{},
		&model.BillDetail{},
		&model.Bill{},
	)
	util.PanicError(err)
}
