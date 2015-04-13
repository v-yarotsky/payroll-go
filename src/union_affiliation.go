package payroll

import "errors"

type UnionAffiliation struct {
	MemberID      int
	MonthlyCharge float64
	charges       map[int]*ServiceCharge
}

func NewUnionAffiliation(memberId int, monthlyCharge float64) *UnionAffiliation {
	return &UnionAffiliation{
		MemberID:      memberId,
		MonthlyCharge: monthlyCharge,
		charges:       make(map[int]*ServiceCharge),
	}
}

func (a *UnionAffiliation) CalculateDeductions(pc *Paycheck) float64 {
	return 1.0
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
