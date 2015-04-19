package payroll

import "errors"
import "payroll/domain"

type ChangeEmployeeClassificationTransaction interface {
	Execute() error
	GetClassification() domain.PaymentClassification
	GetSchedule() domain.PaymentSchedule
}

type BaseChangeClassificationTransaction struct {
	EmployeeID int
}

func (t *BaseChangeClassificationTransaction) Execute(tr ChangeEmployeeClassificationTransaction) error {
	e := GpayrollDatabase.GetEmployee(t.EmployeeID)
	if e == nil {
		return errors.New("employee not found")
	}

	e.PaymentClassification = tr.GetClassification()
	e.PaymentSchedule = tr.GetSchedule()

	return nil
}
