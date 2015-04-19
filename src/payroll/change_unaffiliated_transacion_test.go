package payroll

import "testing"

func TestChangeUnaffiliatedTransaction(t *testing.T) {
	empId := 3
	memberId := 7734
	addTr := NewAddHourlyEmployeeTransaction(empId, "Bob", "Home", 12.5)
	addTr.Execute()

	e := GpayrollDatabase.GetEmployee(empId)
	if e == nil {
		t.Fatalf("expected employee to be in the database")
	}

	unionTr := NewChangeMemberTransaction(empId, memberId, 99.95)
	unionTr.Execute()

	if _, ok := e.Affiliation.(*UnionAffiliation); !ok {
		t.Fatal("expected employee to be union member")
	}

	if m, _ := GpayrollDatabase.GetUnionMember(memberId); m == nil {
		t.Fatal("expected employee to be registered as union member")
	}

	tr := NewChangeUnaffiliatedTransaction(empId)
	tr.Execute()

	if _, ok := e.Affiliation.(*NoAffiliation); !ok {
		t.Fatalf("expected employee not to be a union member, got %v", e.Affiliation)
	}

	if m, _ := GpayrollDatabase.GetUnionMember(memberId); m != nil {
		t.Fatalf("expected employee not to be registered as a union member")
	}
}
