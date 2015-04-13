package payroll

import "errors"
import "time"

type UnionAffiliation struct {
	MemberID int
	Dues     float64
	charges  map[int]*ServiceCharge
}

func NewUnionAffiliation(memberId int, dues float64) *UnionAffiliation {
	return &UnionAffiliation{
		MemberID: memberId,
		Dues:     dues,
		charges:  make(map[int]*ServiceCharge),
	}
}

func (a *UnionAffiliation) CalculateDeductions(pc *Paycheck) float64 {
	fridays := a.numberOfFridaysInPayPeriod(pc.PayPeriodStartDate, pc.PayDate)
	return float64(fridays) * a.Dues
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

func (a *UnionAffiliation) GetServiceCharge(date int) (*ServiceCharge, error) {
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
