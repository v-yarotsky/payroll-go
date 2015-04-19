package payroll

import "testing"

func TestChangeAddressTransaction(t *testing.T) {
	empId := 43

	addTr := NewAddHourlyEmployeeTransaction(empId, "Bill", "Home", 12.50)
	addTr.Execute()

	tr := NewChangeAddressTransaction(empId, "Alaska")
	tr.Execute()

	e := GpayrollDatabase.GetEmployee(empId)
	if e == nil {
		t.Fatalf("expected employee to be in the database")
	}

	if e.Address != "Alaska" {
		t.Fatalf("expected employee address to have been changed to Alaska")
	}
}
