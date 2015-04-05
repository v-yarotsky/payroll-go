package payroll

import "errors"

type HourlyClassification struct {
	HourlyRate float64
}

func (c *HourlyClassification) GetTimeCard(date int) (*TimeCard, error) {
	return nil, errors.New("card not found")
}

func (c *HourlyClassification) AddTimeCard(tc *TimeCard) {
}
