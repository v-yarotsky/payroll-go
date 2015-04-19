package payroll

import "time"
import "errors"
import "payroll/domain"

type paydayTransaction struct {
	Date      time.Time
	paychecks map[int]*domain.Paycheck
}

func NewPaydayTransaction(date time.Time) *paydayTransaction {
	return &paydayTransaction{date, make(map[int]*domain.Paycheck)}
}

func (t *paydayTransaction) Execute() error {
	employees := GpayrollDatabase.GetAllEmployees()
	for _, employee := range employees {
		if employee.IsPayDate(t.Date) {
			pc := domain.NewPaycheck(employee.GetPayPeriodStartDate(t.Date), t.Date)
			t.paychecks[employee.ID] = pc
			employee.Payday(pc)
		}
	}
	return nil
}

func (t *paydayTransaction) GetPaycheck(empId int) (*domain.Paycheck, error) {
	pc, ok := t.paychecks[empId]
	if !ok {
		return nil, errors.New("no paycheck for employee " + string(empId))
	}
	return pc, nil
}
