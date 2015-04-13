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

func (c *CommissionedClassification) CalculatePay(pc *Paycheck) float64 {
	totalSales := 0.0
	for _, receipt := range c.salesReceipts {
		if c.isInPayPeriod(receipt, pc.PayDate) {
			totalSales = totalSales + receipt.Amount
		}
	}
	return c.Salary + c.CommissionRate*totalSales
}

func (c *CommissionedClassification) isInPayPeriod(r *SalesReceipt, payPeriod time.Time) bool {
	payPeriodEndDate := payPeriod
	payPeriodStartDate := payPeriod.Add(-13 * 24 * time.Hour)
	receiptDate := r.Date
	return (receiptDate.Equal(payPeriodStartDate) || receiptDate.After(payPeriodStartDate)) &&
		(receiptDate.Equal(payPeriodEndDate) || receiptDate.Before(payPeriodEndDate))
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
