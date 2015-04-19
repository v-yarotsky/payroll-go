package payroll

import "payroll/domain"

type AddEmployeeTransaction interface {
	GetClassification() domain.PaymentClassification
	GetSchedule() domain.PaymentSchedule
	GetMethod() domain.PaymentMethod
	Execute() error
}

type BasicAddEmployeeTransaction struct {
	EmployeeID      int
	EmployeeName    string
	EmployeeAddress string
}

func (t BasicAddEmployeeTransaction) Execute(tr AddEmployeeTransaction) error {
	pc := tr.GetClassification()
	ps := tr.GetSchedule()
	pm := tr.GetMethod()
	aff := NewNoAffiliation()
	e := &domain.Employee{t.EmployeeID, t.EmployeeName, t.EmployeeAddress, ps, pc, pm, aff}
	GpayrollDatabase.AddEmployee(t.EmployeeID, e)
	return nil
}

func (t BasicAddEmployeeTransaction) GetClassification() domain.PaymentClassification { return nil }
func (t BasicAddEmployeeTransaction) GetSchedule() domain.PaymentSchedule             { return nil }
func (t BasicAddEmployeeTransaction) GetMethod() domain.PaymentMethod                 { return nil }
