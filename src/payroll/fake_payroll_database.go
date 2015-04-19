package payroll

import "errors"
import "payroll/domain"

var GpayrollDatabase = NewFakePayrollDatabase()

type PayrollDatabase interface {
	GetAllEmployees() []*domain.Employee
	GetEmployee(int) *domain.Employee
	AddEmployee(int, *domain.Employee)
	DeleteEmployee(int) error
	GetUnionMember(memberId int) (*domain.Employee, error)
	AddUnionMember(memberId int, employee *domain.Employee) error
	RemoveUnionMember(memberId int) error
	Clear()
}

type fakePayrollDatabase struct {
	employees    map[int]*domain.Employee
	unionMembers map[int]*domain.Employee
}

func NewFakePayrollDatabase() *fakePayrollDatabase {
	return &fakePayrollDatabase{
		employees:    make(map[int]*domain.Employee),
		unionMembers: make(map[int]*domain.Employee),
	}
}

func (db *fakePayrollDatabase) GetAllEmployees() []*domain.Employee {
	employees := make([]*domain.Employee, 0, len(db.employees))
	for _, employee := range db.employees {
		employees = append(employees, employee)
	}
	return employees
}

func (db *fakePayrollDatabase) GetEmployee(empId int) *domain.Employee {
	return db.employees[empId]
}

func (db *fakePayrollDatabase) AddEmployee(empId int, emp *domain.Employee) {
	db.employees[empId] = emp
}

func (db *fakePayrollDatabase) DeleteEmployee(empId int) error {
	db.employees[empId] = nil
	return nil
}

func (db *fakePayrollDatabase) GetUnionMember(memberId int) (*domain.Employee, error) {
	e, ok := db.unionMembers[memberId]
	if !ok {
		return nil, errors.New("union member not found")
	}
	return e, nil
}

func (db *fakePayrollDatabase) AddUnionMember(memberId int, e *domain.Employee) error {
	db.unionMembers[memberId] = e
	return nil
}

func (db *fakePayrollDatabase) RemoveUnionMember(memberId int) error {
	db.unionMembers[memberId] = nil
	return nil
}

func (db *fakePayrollDatabase) Clear() {
	db.employees = make(map[int]*domain.Employee)
}
