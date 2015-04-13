package payroll

import "testing"

func TestServiceChargeTransaction(t *testing.T) {
	empId := 2

	addTr := NewAddHourlyEmployeeTransaction(empId, "Bill", "Home", 15.25)
	addTr.Execute()

	e := GpayrollDatabase.GetEmployee(empId)
	if e == nil {
		t.Fatalf("employee must have been in database")
	}

	memberId := 86
	af := NewUnionAffiliation(memberId, 12.5)
	e.Affiliation = af

	GpayrollDatabase.AddUnionMember(memberId, e)

	tr := NewServiceChargeTransaction(memberId, parseDate("2001-Nov-01"), 12.95)
	_ = tr.Execute()

	sc, err := af.GetServiceCharge(parseDate("2001-Nov-01"))
	if err != nil {
		t.Fatalf("Expected service charge to be there, got err %v", err)
	}

	if !floatEquals(sc.Amount, 12.95) {
		t.Fatalf("Expected service charge to amount to %v, got %v", 12.95, sc.Amount)
	}
}
