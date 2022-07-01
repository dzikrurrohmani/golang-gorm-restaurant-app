package usecase

import (
	"errors"
	"fmt"
	model "livecode-gorm-wmb/models"
	repository "livecode-gorm-wmb/repositories"
)

type MemberActivationUseCase interface {
	ActivateMember(phoneNumber string) (model.Customer, error)
}

type memberActivationUseCase struct {
	custRepo     repository.CustomerRepository
	discountRepo repository.DiscountRepository
}

func (m *memberActivationUseCase) ActivateMember(phoneNumber string) (model.Customer, error) {
	by := map[string]interface{}{"mobile_phone_no": phoneNumber}
	customerSlice, err := m.custRepo.FindBy(by)
	if err != nil {
		fmt.Println("Data tidak ditemukan.")
		return model.Customer{}, err
	}
	customerSelected := customerSlice[0]
	if customerSelected.IsMember {
		fmt.Println("Data tidak ditemukan.")
		return customerSelected, errors.New("nomor telepon tersebut sudah terdaftar sebagai member")
	}
	err = m.custRepo.UpdateBy(&customerSelected, map[string]interface{}{"is_member": true})
	if err != nil {
		fmt.Println("Aktivasi member gagal dilakukan.")
		return model.Customer{}, err
	}
	fmt.Println("Aktivasi member berhasil.")
	// Belum selesai tambah discount
	// discountSlice, _ := m.discountRepo.FindBy(map[string]interface{}{"id":2})
	// m.custRepo.UpdateAssociation(&customerSelected, "")
	customerSlice, err = m.custRepo.FindBy(by)
	if err != nil {
		fmt.Println("Data tidak ditemukan.")
		return model.Customer{}, err
	}
	customerReturn := customerSlice[0]
	return customerReturn, nil
}

func NewMemberActivationUseCase(custRepo repository.CustomerRepository, discountRepo repository.DiscountRepository) MemberActivationUseCase {
	usecase := new(memberActivationUseCase)
	usecase.custRepo = custRepo
	usecase.discountRepo = discountRepo
	return usecase
}
