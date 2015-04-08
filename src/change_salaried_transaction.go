package payroll

type changeSalariedTransaction struct {
	BaseChangeClassificationTransaction
	Salary float64
}

func NewChangeSalariedTransaction(empId int, salary float64) *changeSalariedTransaction {
	return &changeSalariedTransaction{BaseChangeClassificationTransaction{empId}, salary}
}

func (t *changeSalariedTransaction) Execute() error {
	return t.BaseChangeClassificationTransaction.Execute(t)
}

func (t *changeSalariedTransaction) GetClassification() PaymentClassification {
	return NewSalariedClassification(t.Salary)
}

func (t *changeSalariedTransaction) GetSchedule() PaymentSchedule {
	return &MonthlySchedule{}
}
