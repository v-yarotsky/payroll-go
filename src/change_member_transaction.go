package payroll

type changeMemberTransaction struct {
	BaseChangeAffiliationTransaction
	MemberID      int
	MonthlyCharge float64
}

func NewChangeMemberTransaction(empId int, memberId int, monthlyCharge float64) *changeMemberTransaction {
	return &changeMemberTransaction{BaseChangeAffiliationTransaction{empId}, memberId, monthlyCharge}
}

func (t *changeMemberTransaction) GetAffiliation() Affiliation {
	return NewUnionAffiliation(t.MemberID, t.MonthlyCharge)
}

func (t *changeMemberTransaction) RecordMembership(e *Employee) error {
	return GpayrollDatabase.AddUnionMember(t.MemberID, e)
}

func (t *changeMemberTransaction) Execute() error {
	return t.BaseChangeAffiliationTransaction.Execute(t)
}
