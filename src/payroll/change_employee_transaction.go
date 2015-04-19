package payroll

import "errors"
import "payroll/domain"

type ChangeEmployeeTransaction interface {
	Execute() error
	Change(*domain.Employee) error
}

type BaseChangeEmployeeTransaction struct {
	EmployeeID int
}

func (t *BaseChangeEmployeeTransaction) Execute(tr ChangeEmployeeTransaction) error {
	e := GpayrollDatabase.GetEmployee(t.EmployeeID)
	if e == nil {
		return errors.New("employee not found")
	}
	return tr.Change(e)
}
