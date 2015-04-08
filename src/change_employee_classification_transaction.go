package payroll

import "errors"

type ChangeEmployeeClassificationTransaction interface {
	Execute() error
	GetClassification() PaymentClassification
	GetSchedule() PaymentSchedule
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
