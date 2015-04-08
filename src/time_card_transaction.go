package payroll

import "errors"

type timeCardTransaction struct {
	Date       int
	Hours      float64
	EmployeeID int
}

func NewTimeCardTransaction(date int, hours float64, empId int) timeCardTransaction {
	return timeCardTransaction{date, hours, empId}
}

func (t *timeCardTransaction) Execute() error {
	e := GpayrollDatabase.GetEmployee(t.EmployeeID)
	if e == nil {
		return errors.New("no such employee")
	}

	pc, ok := e.PaymentClassification.(*HourlyClassification)

	if !ok {
		return errors.New("tried to add time card to non-hourly employee")
	}

	tc := &TimeCard{t.Date, t.Hours}
	pc.AddTimeCard(tc)
	return nil
}
