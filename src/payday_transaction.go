package payroll

import "time"
import "errors"

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
	return DateIsBetween(date, p.PayPeriodStartDate, p.PayPeriodEndDate)
}

type paydayTransaction struct {
	Date      time.Time
	paychecks map[int]*Paycheck
}

func NewPaydayTransaction(date time.Time) *paydayTransaction {
	return &paydayTransaction{date, make(map[int]*Paycheck)}
}

func (t *paydayTransaction) Execute() error {
	employees := GpayrollDatabase.GetAllEmployees()
	for _, employee := range employees {
		if employee.IsPayDate(t.Date) {
			pc := NewPaycheck(employee.GetPayPeriodStartDate(t.Date), t.Date)
			t.paychecks[employee.ID] = pc
			employee.Payday(pc)
		}
	}
	return nil
}

func (t *paydayTransaction) GetPaycheck(empId int) (*Paycheck, error) {
	pc, ok := t.paychecks[empId]
	if !ok {
		return nil, errors.New("no paycheck for employee " + string(empId))
	}
	return pc, nil
}
