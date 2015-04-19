package domain

import "time"
import "errors"
import "payroll/util"

type Paycheck struct {
	PayPeriodStartDate time.Time
	PayPeriodEndDate   time.Time
	GrossPay           float64
	Deductions         float64
	NetPay             float64
	fields             map[string]string
}

func NewPaycheck(payPeriodStartDate, payPeriodEndDate time.Time) *Paycheck {
	return &Paycheck{
		PayPeriodStartDate: payPeriodStartDate,
		PayPeriodEndDate:   payPeriodEndDate,
		fields:             make(map[string]string),
	}
}

func (p *Paycheck) GetField(field string) (string, error) {
	if fieldValue, ok := p.fields[field]; ok {
		return fieldValue, nil
	}
	return "", errors.New("field not found")
}

func (p *Paycheck) SetField(fieldName, fieldValue string) {
	p.fields[fieldName] = fieldValue
}

func (p *Paycheck) IsInPayPeriod(date time.Time) bool {
	return util.DateIsBetween(date, p.PayPeriodStartDate, p.PayPeriodEndDate)
}
