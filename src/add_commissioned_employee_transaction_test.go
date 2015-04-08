package payroll

import (
	"testing"
)

func TestAddCommissionedEmployee(t *testing.T) {
	defer GpayrollDatabase.Clear()

	empId := 1
	tr := NewAddCommissionedEmployeeTransaction(empId, "Eric", "Kentucky", 1000.00, 1.00)
	tr.Execute()

	e := GpayrollDatabase.GetEmployee(empId)

	if e.Name != "Eric" {
		t.Fatalf("expected employee name to be %v, got %v", "Eric", e.Name)
	}

	pc, ok := e.PaymentClassification.(*CommissionedClassification)
	if !ok {
		t.Fatalf("expected employee to have commissioned classification, got %x", e.PaymentClassification)
	}

	if !floatEquals(pc.Salary, 1000.00) {
		t.Fatalf("expected salary to be %v, got %v", 1000.00, pc.Salary)
	}

	if !floatEquals(pc.CommissionRate, 1.00) {
		t.Fatalf("expected commission rate to be %v, got %v", 1.00, pc.CommissionRate)
	}

	if _, ok := e.PaymentSchedule.(BiweeklySchedule); !ok {
		t.Fatalf("expected employee to have bi-weekly payments schedule, got %x", e.PaymentSchedule)
	}

	if _, ok := e.PaymentMethod.(HoldMethod); !ok {
		t.Fatalf("expected employee to have 'hold' payments method, got %v", e.PaymentMethod)
	}

}
