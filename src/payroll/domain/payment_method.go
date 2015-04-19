package domain

type PaymentMethod interface {
	Pay(*Paycheck)
}
