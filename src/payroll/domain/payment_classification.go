package domain

type PaymentClassification interface {
	CalculatePay(*Paycheck) float64
}
