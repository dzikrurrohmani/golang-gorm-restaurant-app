package command

import (
	model "livecode-gorm-wmb/models"
	repository "livecode-gorm-wmb/repositories"
	usecase "livecode-gorm-wmb/usecases/customer"
	util "livecode-gorm-wmb/utils"
	"log"
)

func RegistCustomer(customerRepo repository.CustomerRepository) {
	// Create (C)
	customerRegistrationUseCase := usecase.NewCustomerRegistrationUseCase(customerRepo)
	customer01 := model.Customer{
		CustomerName: "Anak Pertama",
		MobilePhoneNo: "011111111111",
		
	}
	customer02 := model.Customer{
		CustomerName: "Anak Kedua",
		MobilePhoneNo: "022222222222",
	}
	customer03 := model.Customer{
		CustomerName: "Anak Ketiga",
		MobilePhoneNo: "033333333333",
	}
	err := customerRegistrationUseCase.CreateCustomer([]*model.Customer{&customer01, &customer02, &customer03})
	util.RaiseError(err)
}

func ActivateMember(customerRepo repository.CustomerRepository, discountRepo repository.DiscountRepository) {
	// Update (U)
	memberActivationUseCase := usecase.NewMemberActivationUseCase(customerRepo, discountRepo)
	newMember, err := memberActivationUseCase.ActivateMember("011111111111")
	util.RaiseError(err)
	log.Println(newMember.ToString())
}
