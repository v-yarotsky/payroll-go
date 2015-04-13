package payroll

import "testing"

func TestTimeCardTransaction(t *testing.T) {
	empId := 2

	addTr := NewAddHourlyEmployeeTransaction(empId, "Bill", "Home", 15.25)
	addTr.Execute()

	tcTr := NewTimeCardTransaction(parseDate("2001-Oct-31"), 8.0, empId)
	_ = tcTr.Execute()

	e := GpayrollDatabase.GetEmployee(empId)
	if e == nil {
		t.Fatalf("employee must have been in database")
	}

	pc, ok := e.PaymentClassification.(*HourlyClassification)

	if !ok {
		t.Fatalf("expected hourly payment classification")
	}

	tc, err := pc.GetTimeCard(parseDate("2001-Oct-31"))

	if err != nil {
		t.Fatalf("expected to find a time card, got error %v", err)
	}

	if !floatEquals(tc.Hours, 8.0) {
		t.Fatalf("expected time card with %v hours, got %v hours", 8.0, tc.Hours)
	}
}
