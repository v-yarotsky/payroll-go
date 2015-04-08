package payroll

import "errors"

var GpayrollDatabase = NewFakePayrollDatabase()

type PayrollDatabase interface {
	GetEmployee(int) *Employee
	AddEmployee(int, *Employee)
	DeleteEmployee(int)
	GetUnionMember(memberId int) (*Employee, error)
	AddUnionMember(memberId int, employee *Employee) error
	Clear()
}

type fakePayrollDatabase struct {
	employees    map[int]*Employee
	unionMembers map[int]*Employee
}

func NewFakePayrollDatabase() *fakePayrollDatabase {
	return &fakePayrollDatabase{
		employees:    make(map[int]*Employee),
		unionMembers: make(map[int]*Employee),
	}
}

func (db *fakePayrollDatabase) GetEmployee(empId int) *Employee {
	return db.employees[empId]
}

func (db *fakePayrollDatabase) AddEmployee(empId int, emp *Employee) {
	db.employees[empId] = emp
}

func (db *fakePayrollDatabase) DeleteEmployee(empId int) {
	db.employees[empId] = nil
}

func (db *fakePayrollDatabase) GetUnionMember(memberId int) (*Employee, error) {
	e, ok := db.unionMembers[memberId]
	if !ok {
		return nil, errors.New("union member not found")
	}
	return e, nil
}

func (db *fakePayrollDatabase) AddUnionMember(memberId int, e *Employee) error {
	db.unionMembers[memberId] = e
	return nil
}

func (db *fakePayrollDatabase) Clear() {
	db.employees = make(map[int]*Employee)
}
