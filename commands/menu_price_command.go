package command

import (
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
		MenuID: 5,
	}
	menuPrice02 := model.MenuPrice{
		Price: 5000,
		MenuID: 6,
	}
	menuPrice03 := model.MenuPrice{
		Price: 10000,
		MenuID: 7,
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
	menuPricewithID, err := ReadMenuPriceUseCase.ReadMenuPriceById(3)
	util.RaiseError(err)
	log.Println(menuPricewithID)

	// Proses delete dilakukan
	DeleteMenuPriceUseCase := usecase.NewDeleteMenuPriceUseCase(menuPriceRepo)
	DeleteMenuPriceUseCase.DeleteMenuPrice(&menuPricewithID)

	// Dicetak kembali untuk dicek apakah sudah berubah
	menuPricewithID, err = ReadMenuPriceUseCase.ReadMenuPriceById(3)
	util.RaiseError(err)
	log.Println(menuPricewithID)
}
