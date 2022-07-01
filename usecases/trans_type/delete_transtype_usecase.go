package usecase

import (
	model "livecode-gorm-wmb/models"
	repository "livecode-gorm-wmb/repositories"
)

type DeleteTransTypeUseCase interface {
	DeleteTransType(transType *model.TransType) error
}

type deleteTransTypeUseCase struct {
	transTypeRepo repository.TransTypeRepository
}

func (m *deleteTransTypeUseCase) DeleteTransType(transType *model.TransType) error {
	return m.transTypeRepo.Delete(transType)
}

func NewDeleteTransTypeUseCase(transTypeRepo repository.TransTypeRepository) DeleteTransTypeUseCase {
	usecase := new(deleteTransTypeUseCase)
	usecase.transTypeRepo = transTypeRepo
	return usecase
}
