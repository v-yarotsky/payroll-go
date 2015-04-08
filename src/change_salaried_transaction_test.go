package payroll

import "testing"

func TestChangeSalariedTransactionTest(t *testing.T) {
	empId := 7

	addTr := NewAddCommissionedEmployeeTransaction(empId, "Bill", "Home", 1000.0, 0.001)
	addTr.Execute()

	tr := NewChangeSalariedTransaction(empId, 1000.0)
	tr.Execute()

	e := GpayrollDatabase.GetEmployee(empId)
	if e == nil {
		t.Fatalf("expected employee to be in the database")
	}

	pc, ok := e.PaymentClassification.(*SalariedClassification)
	if !ok {
		t.Fatalf("expected employee to become salaried employee, but became %x", e.PaymentClassification)
	}

	if !floatEquals(pc.Salary, 1000.0) {
		t.Fatalf("expected salary to be 1000.0, got %v", pc.Salary)
	}

	_, ok = e.PaymentSchedule.(*MonthlySchedule)
	if !ok {
		t.Fatalf("expected employee's schedule to change to monthly schedule, got %x", e.PaymentSchedule)
	}
}
