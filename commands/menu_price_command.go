package command

import (
	"database/sql"
	model "livecode-gorm-wmb/models"
	repository "livecode-gorm-wmb/repositories"
	usecase "livecode-gorm-wmb/usecases/menu_price"
	util "livecode-gorm-wmb/utils"
	"log"
)

func MenuPriceCreate(menuPriceRepo repository.MenuPriceRepository) {
	// Create (C)
	CreateMenuPriceUseCase := usecase.NewCreateMenuPriceUseCase(menuPriceRepo)
	menuPrice01 := model.MenuPrice{
		Price: 2000,
	}
	menuPrice02 := model.MenuPrice{
		Price: 5000,
		MenuID: sql.NullInt64{Int64: 2, Valid: true},
	}
	menuPrice03 := model.MenuPrice{
		Price: 10000,
		MenuID: sql.NullInt64{Int64: 3, Valid: true},
	}
	err := CreateMenuPriceUseCase.CreateMenuPrice([]*model.MenuPrice{&menuPrice01, &menuPrice02, &menuPrice03})
	util.RaiseError(err)
}

func MenuPriceRead(menuPriceRepo repository.MenuPriceRepository) {
	// Read (R)
	ReadMenuPriceUseCase := usecase.NewReadMenuPriceUseCase(menuPriceRepo)
	menuPriceAll, err := ReadMenuPriceUseCase.ReadAllMenuPrice()
	util.RaiseError(err)
	log.Println(menuPriceAll)

	menuPricewithID, err := ReadMenuPriceUseCase.ReadMenuPriceById(3)
	util.RaiseError(err)
	log.Println(menuPricewithID)

	menuPricewithName, err := ReadMenuPriceUseCase.ReadMenuPriceBy(map[string]interface{}{"menuPrice_name": "Nasi Uduk"})
	util.RaiseError(err)
	log.Println(menuPricewithName)
}

func MenuPriceUpdate(menuPriceRepo repository.MenuPriceRepository) {
	// Cari menuPrice yang ingin di update
	ReadMenuPriceUseCase := usecase.NewReadMenuPriceUseCase(menuPriceRepo)
	menuPricewithID, err := ReadMenuPriceUseCase.ReadMenuPriceById(3)
	util.RaiseError(err)
	log.Println(menuPricewithID)

	// Misal ganti nama
	UpdateMenuPriceUseCase := usecase.NewUpdateMenuPriceUseCase(menuPriceRepo)
	UpdateMenuPriceUseCase.UpdateMenuPrice(&menuPricewithID, map[string]interface{}{"menuPrice_name": "Nasi Liwet"})

	// Dicetak kembali untuk dicek apakah sudah berubah
	menuPricewithID, err = ReadMenuPriceUseCase.ReadMenuPriceById(3)
	util.RaiseError(err)
	log.Println(menuPricewithID)
}

func MenuPriceDelete(menuPriceRepo repository.MenuPriceRepository) {
	// Cari menuPrice yang ingin di update
	ReadMenuPriceUseCase := usecase.NewReadMenuPriceUseCase(menuPriceRepo)
	menuPricewithID, err := ReadMenuPriceUseCase.ReadMenuPriceById(1)
	util.RaiseError(err)
	log.Println(menuPricewithID)

	// Proses delete dilakukan
	DeleteMenuPriceUseCase := usecase.NewDeleteMenuPriceUseCase(menuPriceRepo)
	DeleteMenuPriceUseCase.DeleteMenuPrice(&menuPricewithID)

	// Dicetak kembali untuk dicek apakah sudah berubah
	menuPricewithID, err = ReadMenuPriceUseCase.ReadMenuPriceById(1)
	util.RaiseError(err)
	log.Println(menuPricewithID)
}
