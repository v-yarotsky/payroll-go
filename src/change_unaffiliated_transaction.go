package payroll

type changeUnaffiliatedTransaction struct {
	BaseChangeAffiliationTransaction
}

func NewChangeUnaffiliatedTransaction(empId int) *changeUnaffiliatedTransaction {
	return &changeUnaffiliatedTransaction{BaseChangeAffiliationTransaction{empId}}
}

func (t *changeUnaffiliatedTransaction) GetAffiliation() Affiliation {
	return NewNoAffiliation()
}

func (t *changeUnaffiliatedTransaction) RecordMembership(e *Employee) error {
	if a, ok := e.Affiliation.(*UnionAffiliation); ok {
		return GpayrollDatabase.RemoveUnionMember(a.MemberID)
	}
	return nil
}

func (t *changeUnaffiliatedTransaction) Execute() error {
	return t.BaseChangeAffiliationTransaction.Execute(t)
}
