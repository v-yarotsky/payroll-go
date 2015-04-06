package payroll

type addHourlyEmployeeTransaction struct {
	BasicAddEmployeeTransaction
	HourlyRate float64
}

func NewAddHourlyEmployeeTransaction(empId int, name string, address string, hourlyRate float64) addHourlyEmployeeTransaction {
	return addHourlyEmployeeTransaction{BasicAddEmployeeTransaction{empId, name, address}, hourlyRate}
}

func (t addHourlyEmployeeTransaction) GetClassification() PaymentClassification {
	return HourlyClassification{EmployeeID: t.EmployeeID, HourlyRate: t.HourlyRate}
}

func (t addHourlyEmployeeTransaction) GetSchedule() PaymentSchedule {
	return WeeklySchedule{}
}

func (t addHourlyEmployeeTransaction) GetMethod() PaymentMethod {
	return HoldMethod{}
}

func (t addHourlyEmployeeTransaction) Execute() {
	t.BasicAddEmployeeTransaction.Execute(t)
}
