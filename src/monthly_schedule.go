package payroll

import "time"

type MonthlySchedule struct {
}

func (s *MonthlySchedule) IsPayDate(date time.Time) bool {
	m1 := date.Month()
	m2 := date.Add(24 * time.Hour).Month()
	return m1 != m2
}

func (s *MonthlySchedule) GetPayPeriodStartDate(payPeriodEndDate time.Time) time.Time {
	return time.Date(payPeriodEndDate.Year(), payPeriodEndDate.Month(), 1, 0, 0, 0, 0, time.UTC)
}
