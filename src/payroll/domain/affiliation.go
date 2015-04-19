package domain

type Affiliation interface {
	CalculateDeductions(*Paycheck) float64
}
