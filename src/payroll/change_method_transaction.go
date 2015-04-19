package payroll

import "errors"
import "payroll/domain"

type ChangeMethodTransaction interface {
	Execute() error
	GetMethod() domain.PaymentMethod
}

type BaseChangeMethodTransaction struct {
	EmployeeID int
}

func (t *BaseChangeMethodTransaction) Execute(tr ChangeMethodTransaction) error {
	e := GpayrollDatabase.GetEmployee(t.EmployeeID)
	if e == nil {
		return errors.New("employee not found")
	}
	e.PaymentMethod = tr.GetMethod()
	return nil
}
