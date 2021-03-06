package payroll

import "errors"
import "time"

type SalesReceipt struct {
	Date   time.Time
	Amount float64
}

type salesReceiptTransaction struct {
	EmployeeID int
	Date       time.Time
	Amount     float64
}

func NewSalesReceiptTransaction(empId int, date time.Time, amount float64) salesReceiptTransaction {
	return salesReceiptTransaction{empId, date, amount}
}

func (t salesReceiptTransaction) Execute() error {
	e := GpayrollDatabase.GetEmployee(t.EmployeeID)
	if e == nil {
		return errors.New("no such employee")
	}

	pc, ok := e.PaymentClassification.(*CommissionedClassification)

	if !ok {
		return errors.New("tried to add sales receipt to non-commissioned employee")
	}

	sr := &SalesReceipt{t.Date, t.Amount}
	pc.AddSalesReceipt(sr)
	return nil
}
