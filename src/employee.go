package payroll

import "time"

type Employee struct {
	ID                    int
	Name                  string
	Address               string
	PaymentSchedule       PaymentSchedule
	PaymentClassification PaymentClassification
	PaymentMethod         PaymentMethod
	Affiliation           Affiliation
}

func (e *Employee) IsPayDate(date time.Time) bool {
	return e.PaymentSchedule.IsPayDate(date)
}

func (e *Employee) GetPayPeriodStartDate(payDate time.Time) time.Time {
	return e.PaymentSchedule.GetPayPeriodStartDate(payDate)
}

func (e *Employee) Payday(pc *Paycheck) {
	grossPay := e.PaymentClassification.CalculatePay(pc)
	deductions := e.Affiliation.CalculateDeductions(pc)
	netPay := grossPay - deductions
	pc.GrossPay = grossPay
	pc.Deductions = deductions
	pc.NetPay = netPay
	e.PaymentMethod.Pay(pc)
}
