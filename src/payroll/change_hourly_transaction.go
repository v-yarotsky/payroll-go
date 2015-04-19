package payroll

import "payroll/domain"

type changeHourlyTransaction struct {
	BaseChangeClassificationTransaction
	HourlyRate float64
}

func NewChangeHourlyTransaction(empId int, hourlyRate float64) *changeHourlyTransaction {
	return &changeHourlyTransaction{BaseChangeClassificationTransaction{empId}, hourlyRate}
}

func (t *changeHourlyTransaction) Execute() error {
	return t.BaseChangeClassificationTransaction.Execute(t)
}

func (t *changeHourlyTransaction) GetClassification() domain.PaymentClassification {
	return NewHourlyClassification(t.HourlyRate)
}

func (t *changeHourlyTransaction) GetSchedule() domain.PaymentSchedule {
	return &WeeklySchedule{}
}
