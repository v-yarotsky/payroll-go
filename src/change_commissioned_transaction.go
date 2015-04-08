package payroll

type changeCommissionedTransaction struct {
	BaseChangeClassificationTransaction
	Salary         float64
	CommissionRate float64
}

func NewChangeCommissionedTransaction(empId int, salary float64, commissionRate float64) *changeCommissionedTransaction {
	return &changeCommissionedTransaction{BaseChangeClassificationTransaction{empId}, salary, commissionRate}
}

func (t *changeCommissionedTransaction) Execute() error {
	return t.BaseChangeClassificationTransaction.Execute(t)
}

func (t *changeCommissionedTransaction) GetClassification() PaymentClassification {
	return NewCommissionedClassification(t.Salary, t.CommissionRate)
}

func (t *changeCommissionedTransaction) GetSchedule() PaymentSchedule {
	return &BiweeklySchedule{}
}
