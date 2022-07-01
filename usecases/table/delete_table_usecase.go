package usecase

import (
	model "livecode-gorm-wmb/models"
	repository "livecode-gorm-wmb/repositories"
)

type DeleteTableUseCase interface {
	DeleteTable(table *model.Table) error
}

type deleteTableUseCase struct {
	tableRepo repository.TableRepository
}

func (m *deleteTableUseCase) DeleteTable(table *model.Table) error {
	return m.tableRepo.Delete(table)
}

func NewDeleteTableUseCase(tableRepo repository.TableRepository) DeleteTableUseCase {
	usecase := new(deleteTableUseCase)
	usecase.tableRepo = tableRepo
	return usecase
}
