package payroll

type addSalariedEmployeeTransaction struct {
	BasicAddEmployeeTransaction
	Salary float64
}

func NewAddSalariedEmployeeTransaction(empId int, name string, address string, salary float64) addSalariedEmployeeTransaction {
	return addSalariedEmployeeTransaction{BasicAddEmployeeTransaction{empId, name, address}, salary}
}

func (t addSalariedEmployeeTransaction) GetClassification() PaymentClassification {
	return NewSalariedClassification(t.Salary)
}

func (t addSalariedEmployeeTransaction) GetSchedule() PaymentSchedule {
	return &MonthlySchedule{}
}

func (t addSalariedEmployeeTransaction) GetMethod() PaymentMethod {
	return &HoldMethod{}
}

func (t addSalariedEmployeeTransaction) Execute() {
	t.BasicAddEmployeeTransaction.Execute(t)
}
