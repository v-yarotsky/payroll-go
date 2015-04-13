package payroll

import "time"

type PaymentSchedule interface {
	IsPayDate(time.Time) bool
	GetPayPeriodStartDate(time.Time) time.Time
}
