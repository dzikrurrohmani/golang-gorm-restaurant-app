package usecase

import (
	model "livecode-gorm-wmb/models"
	repository "livecode-gorm-wmb/repositories"
)

type UpdateTableUseCase interface {
	UpdateTable(table *model.Table,  by map[string]interface{}) error
}

type updateTableUseCase struct {
	tableRepo repository.TableRepository
}

func (m *updateTableUseCase) UpdateTable(table *model.Table,  by map[string]interface{}) error {
	return m.tableRepo.UpdateBy(table, by)
}

func NewUpdateTableUseCase(tableRepo repository.TableRepository) UpdateTableUseCase {
	usecase := new(updateTableUseCase)
	usecase.tableRepo = tableRepo
	return usecase
}
