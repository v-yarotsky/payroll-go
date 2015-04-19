package payroll

import "payroll/domain"

type addCommissionedEmployeeTransaction struct {
	BasicAddEmployeeTransaction
	Salary         float64
	CommissionRate float64
}

func NewAddCommissionedEmployeeTransaction(empId int, name string, address string, salary float64, commissionRate float64) addCommissionedEmployeeTransaction {
	return addCommissionedEmployeeTransaction{
		BasicAddEmployeeTransaction{empId, name, address},
		salary, commissionRate}
}

func (t addCommissionedEmployeeTransaction) GetClassification() domain.PaymentClassification {
	return NewCommissionedClassification(t.Salary, t.CommissionRate)
}

func (t addCommissionedEmployeeTransaction) GetSchedule() domain.PaymentSchedule {
	return &BiweeklySchedule{}
}

func (t addCommissionedEmployeeTransaction) GetMethod() domain.PaymentMethod {
	return &HoldMethod{}
}

func (t addCommissionedEmployeeTransaction) Execute() error {
	return t.BasicAddEmployeeTransaction.Execute(t)
}
