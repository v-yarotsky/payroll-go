package payroll

import "errors"
import "payroll/domain"

type ChangeAffiliationTransaction interface {
	Execute() error
	GetAffiliation() domain.Affiliation
	RecordMembership(*domain.Employee) error
}

type BaseChangeAffiliationTransaction struct {
	EmployeeID int
}

func (t *BaseChangeAffiliationTransaction) Execute(tr ChangeAffiliationTransaction) error {
	e := GpayrollDatabase.GetEmployee(t.EmployeeID)
	if e == nil {
		return errors.New("employee not found")
	}
	tr.RecordMembership(e)
	e.Affiliation = tr.GetAffiliation()
	return nil
}
