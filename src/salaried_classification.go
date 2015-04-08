package payroll

type SalariedClassification struct {
	Salary float64
}

func NewSalariedClassification(salary float64) *SalariedClassification {
	return &SalariedClassification{salary}
}
