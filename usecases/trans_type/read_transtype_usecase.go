package usecase

import (
	model "livecode-gorm-wmb/models"
	repository "livecode-gorm-wmb/repositories"
)

type ReadTransTypeUseCase interface {
	ReadAllTransType() ([]model.TransType, error)
	ReadTransTypeById(id uint) (model.TransType, error)
	ReadTransTypeBy(by map[string]interface{}) ([]model.TransType, error)
}

type readTransTypeUseCase struct {
	transTypeRepo repository.TransTypeRepository
}

func (m *readTransTypeUseCase) ReadAllTransType() ([]model.TransType, error) {
	return m.transTypeRepo.FindAll()
}

func (m *readTransTypeUseCase) ReadTransTypeById(id uint) (model.TransType, error) {
	transTypeSlice, err := m.transTypeRepo.FindBy(map[string]interface{}{"id": id})
	if len(transTypeSlice) == 0 {
		return model.TransType{}, err
	}
	return transTypeSlice[0], err
}

func (m *readTransTypeUseCase) ReadTransTypeBy(by map[string]interface{}) ([]model.TransType, error) {
	return m.transTypeRepo.FindBy(by)
}

func NewReadTransTypeUseCase(transTypeRepo repository.TransTypeRepository) ReadTransTypeUseCase {
	usecase := new(readTransTypeUseCase)
	usecase.transTypeRepo = transTypeRepo
	return usecase
}
