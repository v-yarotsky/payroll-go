package payroll

import "payroll/domain"

type addHourlyEmployeeTransaction struct {
	BasicAddEmployeeTransaction
	HourlyRate float64
}

func NewAddHourlyEmployeeTransaction(empId int, name string, address string, hourlyRate float64) addHourlyEmployeeTransaction {
	return addHourlyEmployeeTransaction{BasicAddEmployeeTransaction{empId, name, address}, hourlyRate}
}

func (t addHourlyEmployeeTransaction) GetClassification() domain.PaymentClassification {
	return NewHourlyClassification(t.HourlyRate)
}

func (t addHourlyEmployeeTransaction) GetSchedule() domain.PaymentSchedule {
	return &WeeklySchedule{}
}

func (t addHourlyEmployeeTransaction) GetMethod() domain.PaymentMethod {
	return &HoldMethod{}
}

func (t addHourlyEmployeeTransaction) Execute() error {
	return t.BasicAddEmployeeTransaction.Execute(t)
}
