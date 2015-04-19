package payroll

import "payroll/domain"

type changeMemberTransaction struct {
	BaseChangeAffiliationTransaction
	MemberID int
	Dues     float64
}

func NewChangeMemberTransaction(empId int, memberId int, dues float64) *changeMemberTransaction {
	return &changeMemberTransaction{BaseChangeAffiliationTransaction{empId}, memberId, dues}
}

func (t *changeMemberTransaction) GetAffiliation() domain.Affiliation {
	return NewUnionAffiliation(t.MemberID, t.Dues)
}

func (t *changeMemberTransaction) RecordMembership(e *domain.Employee) error {
	return GpayrollDatabase.AddUnionMember(t.MemberID, e)
}

func (t *changeMemberTransaction) Execute() error {
	return t.BaseChangeAffiliationTransaction.Execute(t)
}
