package payroll

import "time"

var EPSILON float64 = 0.00000001

func floatEquals(a, b float64) bool {
	if (a-b) < EPSILON && (b-a) < EPSILON {
		return true
	}
	return false
}

func parseDate(dateStr string) time.Time {
	t, err := time.Parse("2006-Jan-02", dateStr)
	if err != nil {
		panic(err)
	}
	return t
}
