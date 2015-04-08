package payroll

import "testing"

func TestChangeHourlyTransactionTest(t *testing.T) {
	empId := 7

	addTr := NewAddCommissionedEmployeeTransaction(empId, "Bill", "Home", 1000.0, 0.001)
	addTr.Execute()

	tr := NewChangeHourlyTransaction(empId, 12.5)
	tr.Execute()

	e := GpayrollDatabase.GetEmployee(empId)
	if e == nil {
		t.Fatalf("expected employee to be in the database")
	}

	pc, ok := e.PaymentClassification.(*HourlyClassification)
	if !ok {
		t.Fatalf("expected employee to become hourly employee, but became %x", e.PaymentClassification)
	}

	if !floatEquals(pc.HourlyRate, 12.5) {
		t.Fatalf("expected hourly rate to be 12.5, got %v", pc.HourlyRate)
	}

	_, ok = e.PaymentSchedule.(*WeeklySchedule)
	if !ok {
		t.Fatalf("expected employee's schedule to change to weekly schedule, got %x", e.PaymentSchedule)
	}
}
