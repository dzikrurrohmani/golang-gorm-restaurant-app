package command

import (
	model "livecode-gorm-wmb/models"
	repository "livecode-gorm-wmb/repositories"
	usecase "livecode-gorm-wmb/usecases/table"
	util "livecode-gorm-wmb/utils"
	"log"
)

func TableCreate(tableRepo repository.TableRepository) {
	// Create (C)
	CreateTableUseCase := usecase.NewCreateTableUseCase(tableRepo)
	table01 := model.Table{
		TableDescription: "T001",
	}
	table02 := model.Table{
		TableDescription: "T002",
	}
	table03 := model.Table{
		TableDescription: "T003",
	}
	err := CreateTableUseCase.CreateTable([]*model.Table{&table01, &table02, &table03})
	util.RaiseError(err)
}

func TableRead(tableRepo repository.TableRepository) {
	// Read (R)
	ReadTableUseCase := usecase.NewReadTableUseCase(tableRepo)
	tableAll, err := ReadTableUseCase.ReadAllTable()
	util.RaiseError(err)
	log.Println(tableAll)

	tablewithID, err := ReadTableUseCase.ReadTableById(3)
	util.RaiseError(err)
	log.Println(tablewithID)

	tablewithName, err := ReadTableUseCase.ReadTableBy(map[string]interface{}{"table_name": "Nasi Uduk"})
	util.RaiseError(err)
	log.Println(tablewithName)
}

func TableUpdate(tableRepo repository.TableRepository) {
	// Cari table yang ingin di update
	ReadTableUseCase := usecase.NewReadTableUseCase(tableRepo)
	tablewithID, err := ReadTableUseCase.ReadTableById(3)
	util.RaiseError(err)
	log.Println(tablewithID)

	// Misal ganti nama
	UpdateTableUseCase := usecase.NewUpdateTableUseCase(tableRepo)
	UpdateTableUseCase.UpdateTable(&tablewithID, map[string]interface{}{"table_name": "Nasi Liwet"})

	// Dicetak kembali untuk dicek apakah sudah berubah
	tablewithID, err = ReadTableUseCase.ReadTableById(3)
	util.RaiseError(err)
	log.Println(tablewithID)
}

func TableDelete(tableRepo repository.TableRepository) {
	// Cari table yang ingin di update
	ReadTableUseCase := usecase.NewReadTableUseCase(tableRepo)
	tablewithID, err := ReadTableUseCase.ReadTableById(3)
	util.RaiseError(err)
	log.Println(tablewithID)

	// Proses delete dilakukan
	DeleteTableUseCase := usecase.NewDeleteTableUseCase(tableRepo)
	DeleteTableUseCase.DeleteTable(&tablewithID)

	// Dicetak kembali untuk dicek apakah sudah berubah
	tablewithID, err = ReadTableUseCase.ReadTableById(3)
	util.RaiseError(err)
	log.Println(tablewithID)
}
