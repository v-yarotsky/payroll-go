package payroll

import "testing"

func TestChangeNameTransaction(t *testing.T) {
	empId := 42

	addTr := NewAddHourlyEmployeeTransaction(empId, "Bill", "Home", 12.50)
	addTr.Execute()

	tr := NewChangeNameTransaction(empId, "Bob")
	tr.Execute()

	e := GpayrollDatabase.GetEmployee(empId)
	if e == nil {
		t.Fatalf("expected employee to be in the database")
	}

	if e.Name != "Bob" {
		t.Fatalf("expected employee name to have been changed to Bob")
	}
}
