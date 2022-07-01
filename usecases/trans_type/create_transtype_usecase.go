package usecase

import (
	model "livecode-gorm-wmb/models"
	repository "livecode-gorm-wmb/repositories"
)

type CreateTransTypeUseCase interface {
	CreateTransType(transType []*model.TransType) error
}

type createTransTypeUseCase struct {
	transTypeRepo repository.TransTypeRepository
}

func (m *createTransTypeUseCase) CreateTransType(transType []*model.TransType) error {
	return m.transTypeRepo.Create(transType)
}

func NewCreateTransTypeUseCase(transTypeRepo repository.TransTypeRepository) CreateTransTypeUseCase {
	usecase := new(createTransTypeUseCase)
	usecase.transTypeRepo = transTypeRepo
	return usecase
}
