package payroll

import "testing"

func TestSalesReceiptTransaction(t *testing.T) {
	empId := 2

	addTr := NewAddCommissionedEmployeeTransaction(empId, "Bill", "Home", 1000.0, 0.01)
	addTr.Execute()

	e := GpayrollDatabase.GetEmployee(empId)
	if e == nil {
		t.Fatalf("employee must have been in database")
	}

	tr := NewSalesReceiptTransaction(empId, parseDate("2001-Nov-01"), 100.0)
	_ = tr.Execute()

	pc, ok := e.PaymentClassification.(*CommissionedClassification)
	if !ok {
		t.Fatalf("expected commissioned payment classification")
	}

	sr, err := pc.GetSalesReceipt(parseDate("2001-Nov-01"))
	if err != nil {
		t.Fatalf("Expected sales receipt to be there, got err %v", err)
	}

	if !floatEquals(sr.Amount, 100.0) {
		t.Fatalf("Expected sales receipt to amount to %v, got %v", 12.95, sr.Amount)
	}
}
