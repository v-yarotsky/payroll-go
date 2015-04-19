package payroll

import "payroll/domain"

type changeAddressTransaction struct {
	BaseChangeEmployeeTransaction
	NewAddress string
}

func NewChangeAddressTransaction(empId int, newAddress string) *changeAddressTransaction {
	return &changeAddressTransaction{BaseChangeEmployeeTransaction{empId}, newAddress}
}

func (t *changeAddressTransaction) Execute() error {
	return t.BaseChangeEmployeeTransaction.Execute(t)
}

func (t *changeAddressTransaction) Change(e *domain.Employee) error {
	e.Address = t.NewAddress
	return nil
}
