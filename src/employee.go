package payroll

type Employee struct {
	ID                    int
	Name                  string
	Address               string
	PaymentSchedule       PaymentSchedule
	PaymentClassification PaymentClassification
	PaymentMethod         PaymentMethod
}
