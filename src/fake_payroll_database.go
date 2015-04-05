package payroll

var GpayrollDatabase = NewFakePayrollDatabase()

type PayrollDatabase interface {
	GetEmployee(int) *Employee
	AddEmployee(int, *Employee)
	DeleteEmployee(int)
	Clear()
}

type fakePayrollDatabase struct {
	employees map[int]*Employee
}

func NewFakePayrollDatabase() *fakePayrollDatabase {
	return &fakePayrollDatabase{employees: make(map[int]*Employee)}
}

func (db *fakePayrollDatabase) GetEmployee(empId int) *Employee {
	return db.employees[empId]
}

func (db *fakePayrollDatabase) AddEmployee(empId int, emp *Employee) {
	db.employees[empId] = emp
}

func (db *fakePayrollDatabase) Clear() {
	db.employees = make(map[int]*Employee)
}

func (db *fakePayrollDatabase) DeleteEmployee(empId int) {
	db.employees[empId] = nil
}
