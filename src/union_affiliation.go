package payroll

import "errors"
import "time"

type UnionAffiliation struct {
	MemberID int
	Dues     float64
	charges  map[time.Time]*ServiceCharge
}

func NewUnionAffiliation(memberId int, dues float64) *UnionAffiliation {
	return &UnionAffiliation{
		MemberID: memberId,
		Dues:     dues,
		charges:  make(map[time.Time]*ServiceCharge),
	}
}

func (a *UnionAffiliation) CalculateDeductions(pc *Paycheck) float64 {
	fridays := a.numberOfFridaysInPayPeriod(pc.PayPeriodStartDate, pc.PayPeriodEndDate)
	regularDues := float64(fridays) * a.Dues

	serviceCharges := 0.0
	for _, sc := range a.charges {
		if pc.IsInPayPeriod(sc.Date) {
			serviceCharges = serviceCharges + sc.Amount
		}
	}

	return regularDues + serviceCharges
}

func (a *UnionAffiliation) numberOfFridaysInPayPeriod(start, end time.Time) int {
	fridays := 0
	for start.Before(end) || start.Equal(end) {
		if start.Weekday() == time.Friday {
			fridays = fridays + 1
		}
		start = start.Add(24 * time.Hour)
	}
	return fridays
}

func (a *UnionAffiliation) GetServiceCharge(date time.Time) (*ServiceCharge, error) {
	sc, ok := a.charges[date]
	if !ok {
		return nil, errors.New("service charge not found")
	}
	return sc, nil
}

func (a *UnionAffiliation) AddServiceCharge(charge *ServiceCharge) error {
	a.charges[charge.Date] = charge
	return nil
}
