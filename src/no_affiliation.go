package payroll

type NoAffiliation struct {
}

func NewNoAffiliation() *NoAffiliation {
	return &NoAffiliation{}
}

func (a *NoAffiliation) CalculateDeductions(pc *Paycheck) float64 {
	return 0.0
}
