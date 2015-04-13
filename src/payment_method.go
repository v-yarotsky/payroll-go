package payroll

type PaymentMethod interface {
	Pay(*Paycheck)
}
