package payroll

import "time"

type WeeklySchedule struct {
}

func (s *WeeklySchedule) IsPayDate(date time.Time) bool {
	return date.Weekday() == time.Friday
}
