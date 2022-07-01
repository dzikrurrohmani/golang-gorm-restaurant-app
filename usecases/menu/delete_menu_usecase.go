package usecase

import (
	model "livecode-gorm-wmb/models"
	repository "livecode-gorm-wmb/repositories"
)

type DeleteMenuUseCase interface {
	DeleteMenu(menu *model.Menu) error
}

type deleteMenuUseCase struct {
	menuRepo repository.MenuRepository
}

func (m *deleteMenuUseCase) DeleteMenu(menu *model.Menu) error {
	return m.menuRepo.Delete(menu)
}

func NewDeleteMenuUseCase(menuRepo repository.MenuRepository) DeleteMenuUseCase {
	usecase := new(deleteMenuUseCase)
	usecase.menuRepo = menuRepo
	return usecase
}
