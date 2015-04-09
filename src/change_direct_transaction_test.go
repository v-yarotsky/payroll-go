package payroll

import "testing"

func TestChangeDirectTransaction(t *testing.T) {
	empId := 3
	addTr := NewAddHourlyEmployeeTransaction(empId, "Bob", "Home", 12.5)
	addTr.Execute()

	tr := NewChangeDirectTransaction(empId, "Dream Bank", "999")
	tr.Execute()

	e := GpayrollDatabase.GetEmployee(empId)
	if e == nil {
		t.Fatalf("expected employee to be in the database")
	}

	m, ok := e.PaymentMethod.(*DirectMethod)
	if !ok {
		t.Fatalf("expected employee payment method to be Direct, got %v", e.PaymentMethod)
	}

	if m.Bank != "Dream Bank" {
		t.Fatalf("expected bank to be Dream Bank, got %v", m.Bank)
	}

	if m.Account != "999" {
		t.Fatalf("expected bank account to be 999, got %v", m.Account)
	}
}
