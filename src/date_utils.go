package payroll

import "time"

func DateIsBetween(date, start, end time.Time) bool {
	return (date.Equal(start) || date.After(start)) &&
		(date.Equal(end) || date.Before(end))
}
