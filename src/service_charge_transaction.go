package payroll

import "errors"
import "time"

type ServiceCharge struct {
	Date   time.Time
	Amount float64
}

type serviceChargeTransaction struct {
	MemberID int
	Date     time.Time
	Charge   float64
}

func NewServiceChargeTransaction(memberId int, date time.Time, charge float64) serviceChargeTransaction {
	return serviceChargeTransaction{memberId, date, charge}
}

func (t serviceChargeTransaction) Execute() error {
	e, err := GpayrollDatabase.GetUnionMember(t.MemberID)
	if err != nil {
		return err
	}
	af, ok := e.Affiliation.(*UnionAffiliation)
	if !ok {
		return errors.New("tried to add service charge to a non-union member")
	}
	charge := &ServiceCharge{t.Date, t.Charge}
	return af.AddServiceCharge(charge)
}
