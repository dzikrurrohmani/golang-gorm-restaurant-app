package command

import (
	model "livecode-gorm-wmb/models"
	repository "livecode-gorm-wmb/repositories"
	usecase "livecode-gorm-wmb/usecases/bill"
	util "livecode-gorm-wmb/utils"
	"time"
)

func BillCreate(billRepo repository.BillRepository, tableRepo repository.TableRepository) {
	// Create (C)
	customerOrderUseCase := usecase.NewCustomerOrderUseCase(billRepo, tableRepo)
	billDetail01 := model.BillDetail{
		MenuPriceID: 1,
		Qty:         5,
	}
	billDetail02 := model.BillDetail{
		MenuPriceID: 2,
		Qty:         3,
	}
	bill01 := model.Bill{
		TransDate:   time.Now(),
		CustomerID:  1,
		TableID:     2,
		TransTypeID: "EI",
		BillDetails: []model.BillDetail{
			billDetail01,
			billDetail02,
		},
	}
	err := customerOrderUseCase.CreateOrder(&bill01)
	util.RaiseError(err)
}

func BillPayment(billRepo repository.BillRepository, tableRepo repository.TableRepository) {
	// Update (U)
	customerPaymentUseCase := usecase.NewCustomerPaymentUseCase(billRepo, tableRepo)
	err := customerPaymentUseCase.OrderPayment(1)
	util.RaiseError(err)
}
