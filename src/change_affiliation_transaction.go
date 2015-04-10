package payroll

import "errors"

type ChangeAffiliationTransaction interface {
	Execute() error
	GetAffiliation() Affiliation
	RecordMembership(*Employee) error
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
