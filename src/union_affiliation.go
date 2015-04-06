package payroll

type ServiceCharge struct {
	Date   int
	Amount float64
}

type UnionAffiliation struct {
	MemberID      int
	MonthlyCharge float64
}

func (a UnionAffiliation) AddServiceCharge(date int, charge float64) error {
	sc := &ServiceCharge{date, charge}
	return GpayrollDatabase.AddServiceCharge(a.MemberID, sc)
}

func (a UnionAffiliation) GetServiceCharge(date int) (*ServiceCharge, error) {
	return GpayrollDatabase.GetServiceCharge(a.MemberID, date)
}
