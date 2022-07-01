package usecase

import (
	model "livecode-gorm-wmb/models"
	repository "livecode-gorm-wmb/repositories"
)

type CreateTableUseCase interface {
	CreateTable(table []*model.Table) error
}

type createTableUseCase struct {
	tableRepo repository.TableRepository
}

func (m *createTableUseCase) CreateTable(table []*model.Table) error {
	return m.tableRepo.Create(table)
}

func NewCreateTableUseCase(tableRepo repository.TableRepository) CreateTableUseCase {
	usecase := new(createTableUseCase)
	usecase.tableRepo = tableRepo
	return usecase
}
