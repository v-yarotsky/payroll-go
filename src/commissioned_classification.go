package payroll

import "errors"
import "time"

type CommissionedClassification struct {
	Salary         float64
	CommissionRate float64
	salesReceipts  map[time.Time]*SalesReceipt
}

func NewCommissionedClassification(salary float64, commissionRate float64) *CommissionedClassification {
	return &CommissionedClassification{
		Salary:         salary,
		CommissionRate: commissionRate,
		salesReceipts:  make(map[time.Time]*SalesReceipt),
	}
}

func (c *CommissionedClassification) GetSalesReceipt(date time.Time) (*SalesReceipt, error) {
	sr, ok := c.salesReceipts[date]
	if !ok {
		return nil, errors.New("sales receipt not found")
	}
	return sr, nil
}

func (c *CommissionedClassification) AddSalesReceipt(sr *SalesReceipt) error {
	c.salesReceipts[sr.Date] = sr
	return nil
}
