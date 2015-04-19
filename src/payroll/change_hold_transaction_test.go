package payroll

import "testing"

func TestChangeHoldTransaction(t *testing.T) {
	empId := 3
	addTr := NewAddHourlyEmployeeTransaction(empId, "Bob", "Home", 12.5)
	addTr.Execute()

	tr := NewChangeHoldTransaction(empId)
	tr.Execute()

	e := GpayrollDatabase.GetEmployee(empId)
	if e == nil {
		t.Fatalf("expected employee to be in the database")
	}

	_, ok := e.PaymentMethod.(*HoldMethod)
	if !ok {
		t.Fatalf("expected employee payment method to be Hold, got %v", e.PaymentMethod)
	}
}
