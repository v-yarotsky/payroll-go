package payroll

import "errors"
import "time"
import "math"

type HourlyClassification struct {
	HourlyRate float64
	timeCards  map[time.Time]*TimeCard
}

func NewHourlyClassification(hourlyRate float64) *HourlyClassification {
	return &HourlyClassification{
		HourlyRate: hourlyRate,
		timeCards:  make(map[time.Time]*TimeCard),
	}
}

// TODO: count weekends as overtime?
func (c *HourlyClassification) CalculatePay(pc *Paycheck) float64 {
	pay := 0.0
	for _, tc := range c.timeCards {
		if c.isInPayPeriod(tc, pc.PayDate) {
			pay = pay + c.calculatePayForTimeCard(tc)
		}
	}
	return pay
}

func (c *HourlyClassification) isInPayPeriod(tc *TimeCard, payPeriod time.Time) bool {
	payPeriodEndDate := payPeriod
	payPeriodStartDate := payPeriod.Add(-5 * 24 * time.Hour)
	timeCardDate := tc.Date
	return (timeCardDate.Equal(payPeriodStartDate) || timeCardDate.After(payPeriodStartDate)) &&
		(timeCardDate.Equal(payPeriodEndDate) || timeCardDate.Before(payPeriodEndDate))
}

func (c *HourlyClassification) calculatePayForTimeCard(tc *TimeCard) float64 {
	straightHours := math.Min(tc.Hours, 8.0)
	overtimeHours := math.Max(tc.Hours-8.0, 0)
	return straightHours*c.HourlyRate + overtimeHours*1.5*c.HourlyRate
}

func (c *HourlyClassification) GetTimeCard(date time.Time) (*TimeCard, error) {
	tc, ok := c.timeCards[date]
	if !ok {
		return nil, errors.New("time card not found")
	}
	return tc, nil
}

func (c *HourlyClassification) AddTimeCard(tc *TimeCard) error {
	c.timeCards[tc.Date] = tc
	return nil
}
