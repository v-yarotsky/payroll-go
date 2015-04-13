package payroll

import "errors"
import "time"

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
