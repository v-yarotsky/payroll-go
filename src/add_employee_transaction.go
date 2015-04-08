package payroll

type AddEmployeeTransaction interface {
	GetClassification() PaymentClassification
	GetSchedule() PaymentSchedule
	GetMethod() PaymentMethod
	Execute()
}

type BasicAddEmployeeTransaction struct {
	EmployeeID      int
	EmployeeName    string
	EmployeeAddress string
}

func (t BasicAddEmployeeTransaction) Execute(tr AddEmployeeTransaction) {
	pc := tr.GetClassification()
	ps := tr.GetSchedule()
	pm := tr.GetMethod()
	aff := NewNoAffiliation()
	e := &Employee{t.EmployeeID, t.EmployeeName, t.EmployeeAddress, ps, pc, pm, aff}
	GpayrollDatabase.AddEmployee(t.EmployeeID, e)
}

func (t BasicAddEmployeeTransaction) GetClassification() PaymentClassification { return nil }
func (t BasicAddEmployeeTransaction) GetSchedule() PaymentSchedule             { return nil }
func (t BasicAddEmployeeTransaction) GetMethod() PaymentMethod                 { return nil }
