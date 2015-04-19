package payroll

import "testing"

func TestDeleteEmployee(t *testing.T) {
	defer GpayrollDatabase.Clear()

	empId := 3
	addTr := NewAddCommissionedEmployeeTransaction(empId, "Lance", "Home", 2500.00, 3.2)
	addTr.Execute()

	e := GpayrollDatabase.GetEmployee(empId)
	if e == nil {
		t.Fatalf("employee must have been found, but it was not")
	}

	delTr := NewDeleteEmployeeTransaction(empId)
	delTr.Execute()

	e = GpayrollDatabase.GetEmployee(empId)
	if e != nil {
		t.Fatalf("employee must have gone, but it is still in the database")
	}
}
