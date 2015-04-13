package payroll

type Affiliation interface {
	CalculateDeductions(*Paycheck) float64
}
