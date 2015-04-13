package payroll

import "time"

type MonthlySchedule struct {
}

func (s *MonthlySchedule) IsPayDate(date time.Time) bool {
	m1 := date.Month()
	m2 := date.Add(24 * time.Hour).Month()
	return m1 != m2
}
