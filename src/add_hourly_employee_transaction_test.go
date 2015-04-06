package payroll

import (
	"testing"
)

func TestAddHourlyEmployee(t *testing.T) {
	defer GpayrollDatabase.Clear()

	empId := 1
	tr := NewAddHourlyEmployeeTransaction(empId, "Alice", "Kentucky", 10.00)
	tr.Execute()

	e := GpayrollDatabase.GetEmployee(empId)

	if e.Name != "Alice" {
		t.Fatalf("expected employee name to be %v, got %v", "Alice", e.Name)
	}

	pc, ok := e.PaymentClassification.(HourlyClassification)
	if !ok {
		t.Fatalf("expected employee to have hourly classification, got %v", e.PaymentClassification)
	}

	if !floatEquals(pc.HourlyRate, 10.00) {
		t.Fatalf("expected hourly rate to be %v, got %v", 10.00, pc.HourlyRate)
	}

	if _, ok := e.PaymentSchedule.(WeeklySchedule); !ok {
		t.Fatalf("expected employee to have weekly payments schedule, got %v", e.PaymentSchedule)
	}

	if _, ok := e.PaymentMethod.(HoldMethod); !ok {
		t.Fatalf("expected employee to have 'hold' payments method, got %v", e.PaymentMethod)
	}

}
