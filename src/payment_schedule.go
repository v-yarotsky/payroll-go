package payroll

import "time"

type PaymentSchedule interface {
	IsPayDate(time.Time) bool
}
