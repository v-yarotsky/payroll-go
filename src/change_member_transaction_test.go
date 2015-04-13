package payroll

import "testing"

func TestChangeMemberTransaction(t *testing.T) {
	empId := 3
	memberId := 7734
	addTr := NewAddHourlyEmployeeTransaction(empId, "Bob", "Home", 12.5)
	addTr.Execute()

	tr := NewChangeMemberTransaction(empId, memberId, 99.42)
	tr.Execute()

	e := GpayrollDatabase.GetEmployee(empId)
	if e == nil {
		t.Fatalf("expected employee to be in the database")
	}

	a, ok := e.Affiliation.(*UnionAffiliation)
	if !ok {
		t.Fatalf("expected employee to be a union member, got %v", e.Affiliation)
	}

	if !floatEquals(a.Dues, 99.42) {
		t.Fatalf("expected union dues to be 99.42, got %v", a.Dues)
	}

	if m, _ := GpayrollDatabase.GetUnionMember(memberId); m != e {
		t.Fatalf("expected employee to be registered as a union member")
	}
}
