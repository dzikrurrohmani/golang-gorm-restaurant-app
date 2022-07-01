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

	// err := db.AutoMigrate(
	// 	&model.Menu{},
	// 	&model.MenuPrice{},
	// 	&model.Discount{},
	// 	&model.Customer{},
	// 	&model.TransType{},
	// 	&model.Tble{},
	// 	&model.Bill{},
	// 	&model.BillDetail{},
	// )
	// util.PanicError(err)

	// err := db.AutoMigrate(
	// 	&model.Menu{},
	// )
	// util.PanicError(err)
	// time.Sleep(5*time.Second)
	// err = db.AutoMigrate(
	// 	&model.MenuPrice{},
	// )
	// util.PanicError(err)
	// time.Sleep(5*time.Second)
	// err = db.AutoMigrate(
	// 	&model.Discount{},
	// )
	// util.PanicError(err)
	// time.Sleep(5*time.Second)
	// err = db.AutoMigrate(
	// 	&model.Customer{},
	// )
	// util.PanicError(err)
	// time.Sleep(5*time.Second)
	// err = db.AutoMigrate(
	// 	&model.TransType{},
	// )
	// util.PanicError(err)
	// time.Sleep(5*time.Second)
	// err = db.AutoMigrate(
	// 	&model.Tble{},
	// )
	// util.PanicError(err)
	// time.Sleep(5*time.Second)
	// err = db.AutoMigrate(
	// 	&model.Bill{},
	// )
	// util.PanicError(err)
	// time.Sleep(5*time.Second)
	// err = db.AutoMigrate(
	// 	&model.BillDetail{},
	// )
	// util.PanicError(err)
}
