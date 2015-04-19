package payroll

import "payroll/domain"

type NoAffiliation struct {
}

func NewNoAffiliation() *NoAffiliation {
	return &NoAffiliation{}
}

func (a *NoAffiliation) CalculateDeductions(pc *domain.Paycheck) float64 {
	return 0.0
}
