package payroll

import "testing"

func TestChangeMailTransaction(t *testing.T) {
	empId := 3
	addTr := NewAddHourlyEmployeeTransaction(empId, "Bob", "Home", 12.5)
	addTr.Execute()

	tr := NewChangeMailTransaction(empId, "New Home")
	tr.Execute()

	e := GpayrollDatabase.GetEmployee(empId)
	if e == nil {
		t.Fatalf("expected employee to be in the database")
	}

	m, ok := e.PaymentMethod.(*MailMethod)
	if !ok {
		t.Fatalf("expected employee payment method to be Mail, got %v", e.PaymentMethod)
	}

	if m.Address != "New Home" {
		t.Fatalf("expected address to be New Home, got %v", m.Address)
	}
}
