package repositories

import (
	"fin/common"
	"fin/models"
)

type PaymentMethodRepository struct {
	Database common.Database
}

func (pmr PaymentMethodRepository) Save(description string) {
	pmr.Database.Insert("payment_methods", "description", description)
}

func (pmr PaymentMethodRepository) List() []models.PaymentMethod {
	rows := pmr.Database.Query("payment_methods")
	pms := models.PaymentMethod{}.ParsePaymentMethods(rows)
	return pms
}
