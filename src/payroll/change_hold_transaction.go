package payroll

import "payroll/domain"

type changeHoldTransaction struct {
	BaseChangeMethodTransaction
}

func NewChangeHoldTransaction(empId int) *changeHoldTransaction {
	return &changeHoldTransaction{BaseChangeMethodTransaction{empId}}
}

func (t *changeHoldTransaction) GetMethod() domain.PaymentMethod {
	return &HoldMethod{}
}

func (t *changeHoldTransaction) Execute() error {
	return t.BaseChangeMethodTransaction.Execute(t)
}
