package payroll

import "payroll/domain"

type addSalariedEmployeeTransaction struct {
	BasicAddEmployeeTransaction
	Salary float64
}

func NewAddSalariedEmployeeTransaction(empId int, name string, address string, salary float64) addSalariedEmployeeTransaction {
	return addSalariedEmployeeTransaction{BasicAddEmployeeTransaction{empId, name, address}, salary}
}

func (t addSalariedEmployeeTransaction) GetClassification() domain.PaymentClassification {
	return NewSalariedClassification(t.Salary)
}

func (t addSalariedEmployeeTransaction) GetSchedule() domain.PaymentSchedule {
	return &MonthlySchedule{}
}

func (t addSalariedEmployeeTransaction) GetMethod() domain.PaymentMethod {
	return &HoldMethod{}
}

func (t addSalariedEmployeeTransaction) Execute() error {
	return t.BasicAddEmployeeTransaction.Execute(t)
}
