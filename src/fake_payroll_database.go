package payroll

import "errors"

var GpayrollDatabase = NewFakePayrollDatabase()

type PayrollDatabase interface {
	GetEmployee(int) *Employee
	AddEmployee(int, *Employee)
	DeleteEmployee(int)
	GetTimeCard(empId int, date int) error
	AddTimeCard(empId int, tc *TimeCard) error
	Clear()
}

type fakePayrollDatabase struct {
	employees map[int]*Employee
	timeCards map[int][]*TimeCard
}

func NewFakePayrollDatabase() *fakePayrollDatabase {
	return &fakePayrollDatabase{
		employees: make(map[int]*Employee),
		timeCards: make(map[int][]*TimeCard),
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

func (db *fakePayrollDatabase) GetTimeCard(empId int, date int) (*TimeCard, error) {
	for _, tc := range db.timeCards[empId] {
		if tc.Date == date {
			return tc, nil
		}
	}
	return nil, errors.New("not found")
}

func (db *fakePayrollDatabase) AddTimeCard(empId int, tc *TimeCard) error {
	db.timeCards[empId] = append(db.timeCards[empId], tc)
	return nil
}

func (db *fakePayrollDatabase) Clear() {
	db.employees = make(map[int]*Employee)
}
