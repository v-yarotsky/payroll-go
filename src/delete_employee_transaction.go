package payroll

type deleteEmployeeTransaction struct {
	EmployeeID int
}

func NewDeleteEmployeeTransaction(empId int) deleteEmployeeTransaction {
	return deleteEmployeeTransaction{empId}
}

func (t deleteEmployeeTransaction) Execute() {
	GpayrollDatabase.DeleteEmployee(t.EmployeeID)
}
