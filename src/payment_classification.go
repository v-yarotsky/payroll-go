package payroll

type PaymentClassification interface {
	CalculatePay(*Paycheck) float64
}
