package payroll

type deleteEmployeeTransaction struct {
	EmployeeID int
}

func NewDeleteEmployeeTransaction(empId int) deleteEmployeeTransaction {
	return deleteEmployeeTransaction{empId}
}

func (t deleteEmployeeTransaction) Execute() error {
	return GpayrollDatabase.DeleteEmployee(t.EmployeeID)
}
