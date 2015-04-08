package payroll

import "errors"

type HourlyClassification struct {
	HourlyRate float64
	timeCards  map[int]*TimeCard
}

func NewHourlyClassification(hourlyRate float64) *HourlyClassification {
	return &HourlyClassification{
		HourlyRate: hourlyRate,
		timeCards:  make(map[int]*TimeCard),
	}
}

func (c *HourlyClassification) GetTimeCard(date int) (*TimeCard, error) {
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
