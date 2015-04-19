package payroll

type Transaction interface {
	Execute() error
}
