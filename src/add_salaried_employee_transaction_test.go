package payroll

import (
	"testing"
)

func TestAddSalariedEmployee(t *testing.T) {
	defer GpayrollDatabase.Clear()

	empId := 1
	tr := NewAddSalariedEmployeeTransaction(empId, "Bob", "Home", 1000.00)
	tr.Execute()

	e := GpayrollDatabase.GetEmployee(empId)

	if e.Name != "Bob" {
		t.Fatalf("expected employee name to be %v, got %v", "Bob", e.Name)
	}

	pc, ok := e.PaymentClassification.(*SalariedClassification)
	if !ok {
		t.Fatalf("expected employee to have salaried classification, got %v", e.PaymentClassification)
	}

	if !floatEquals(pc.Salary, 1000.00) {
		t.Fatalf("expected salary to be %v, got %v", 1000.00, pc.Salary)
	}

	if _, ok := e.PaymentSchedule.(MonthlySchedule); !ok {
		t.Fatalf("expected employee to have monthly payments schedule, got %v", e.PaymentSchedule)
	}

	if _, ok := e.PaymentMethod.(HoldMethod); !ok {
		t.Fatalf("expected employee to have 'hold' payments method, got %v", e.PaymentMethod)
	}

}
