package payroll

import "payroll/domain"

type SalariedClassification struct {
	Salary float64
}

func NewSalariedClassification(salary float64) *SalariedClassification {
	return &SalariedClassification{salary}
}

func (c *SalariedClassification) CalculatePay(pc *domain.Paycheck) float64 {
	return c.Salary
}
