package payroll

import "testing"

func TestChangeCommissionedTransactionTest(t *testing.T) {
	empId := 7

	addTr := NewAddHourlyEmployeeTransaction(empId, "Bill", "Home", 12.5)
	addTr.Execute()

	tr := NewChangeCommissionedTransaction(empId, 1000.0, 0.001)
	tr.Execute()

	e := GpayrollDatabase.GetEmployee(empId)
	if e == nil {
		t.Fatalf("expected employee to be in the database")
	}

	pc, ok := e.PaymentClassification.(*CommissionedClassification)
	if !ok {
		t.Fatalf("expected employee to become commissioned employee, but became %x", e.PaymentClassification)
	}

	if !floatEquals(pc.Salary, 1000.0) {
		t.Fatalf("expected salary to be 1000.0, got %v", pc.Salary)
	}

	if !floatEquals(pc.CommissionRate, 0.001) {
		t.Fatalf("expected commission rate to be 0.001, got %v", pc.CommissionRate)
	}

	_, ok = e.PaymentSchedule.(*BiweeklySchedule)
	if !ok {
		t.Fatalf("expected employee's schedule to change to biweekly schedule, got %x", e.PaymentSchedule)
	}
}
