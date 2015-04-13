package payroll

import "time"

type BiweeklySchedule struct {
}

func (s *BiweeklySchedule) IsPayDate(date time.Time) bool {
	isMidMonth := date.Day() == 15 // should rather be 15th or second Friday, whichever comes first
	m1 := date.Month()
	m2 := date.Add(24 * time.Hour).Month()
	isLastDayOfMonth := m1 != m2
	return isMidMonth || isLastDayOfMonth
}
