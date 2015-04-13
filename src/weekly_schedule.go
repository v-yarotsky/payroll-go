package payroll

import "time"

type WeeklySchedule struct {
}

func (s *WeeklySchedule) IsPayDate(date time.Time) bool {
	return date.Weekday() == time.Friday
}

func (s *WeeklySchedule) GetPayPeriodStartDate(payPeriodEndDate time.Time) time.Time {
	return payPeriodEndDate.Add(-5 * 24 * time.Hour)
}
