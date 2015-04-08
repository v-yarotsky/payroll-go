package payroll

import "errors"

type ChangeEmployeeTransaction interface {
	Execute() error
	Change(*Employee) error
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
