package usecase

import (
	model "livecode-gorm-wmb/models"
	repository "livecode-gorm-wmb/repositories"
)

type DeleteMenuPriceUseCase interface {
	DeleteMenuPrice(menu *model.MenuPrice) error
}

type deleteMenuPriceUseCase struct {
	menuPriceRepo repository.MenuPriceRepository
}

func (m *deleteMenuPriceUseCase) DeleteMenuPrice(menu *model.MenuPrice) error {
	return m.menuPriceRepo.Delete(menu)
}

func NewDeleteMenuPriceUseCase(menuPriceRepo repository.MenuPriceRepository) DeleteMenuPriceUseCase {
	usecase := new(deleteMenuPriceUseCase)
	usecase.menuPriceRepo = menuPriceRepo
	return usecase
}
