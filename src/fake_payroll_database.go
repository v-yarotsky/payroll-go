package payroll

import "errors"

var GpayrollDatabase = NewFakePayrollDatabase()

type PayrollDatabase interface {
	GetEmployee(int) *Employee
	AddEmployee(int, *Employee)
	DeleteEmployee(int)
	GetTimeCard(empId int, date int) error
	AddTimeCard(empId int, tc *TimeCard) error
	GetUnionMember(memberId int) (*Employee, error)
	AddUnionMember(memberId int, employee *Employee) error
	AddServiceCharge(memberId int, date int, sc *ServiceCharge) error
	GetServiceCharge(memberId int, date int) (*ServiceCharge, error)
	GetSalesReceipt(empId int, date int) (*SalesReceipt, error)
	AddSalesReceipt(empId int, sr *SalesReceipt) error
	Clear()
}

type fakePayrollDatabase struct {
	employees      map[int]*Employee
	timeCards      map[int][]*TimeCard
	unionMembers   map[int]*Employee
	serviceCharges map[int][]*ServiceCharge
	salesReceipts  map[int][]*SalesReceipt
}

func NewFakePayrollDatabase() *fakePayrollDatabase {
	return &fakePayrollDatabase{
		employees:      make(map[int]*Employee),
		timeCards:      make(map[int][]*TimeCard),
		unionMembers:   make(map[int]*Employee),
		serviceCharges: make(map[int][]*ServiceCharge),
		salesReceipts:  make(map[int][]*SalesReceipt),
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
	return nil, errors.New("time card not found")
}

func (db *fakePayrollDatabase) AddTimeCard(empId int, tc *TimeCard) error {
	db.timeCards[empId] = append(db.timeCards[empId], tc)
	return nil
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

func (db *fakePayrollDatabase) AddServiceCharge(memberId int, sc *ServiceCharge) error {
	db.serviceCharges[memberId] = append(db.serviceCharges[memberId], sc)
	return nil
}

func (db *fakePayrollDatabase) GetServiceCharge(memberId int, date int) (*ServiceCharge, error) {
	for _, sc := range db.serviceCharges[memberId] {
		if sc.Date == date {
			return sc, nil
		}
	}
	return nil, errors.New("service charge not found")
}

func (db *fakePayrollDatabase) GetSalesReceipt(empId int, date int) (*SalesReceipt, error) {
	for _, sr := range db.salesReceipts[empId] {
		if sr.Date == date {
			return sr, nil
		}
	}
	return nil, errors.New("sales receipt not found")
}

func (db *fakePayrollDatabase) AddSalesReceipt(empId int, sr *SalesReceipt) error {
	db.salesReceipts[empId] = append(db.salesReceipts[empId], sr)
	return nil
}

func (db *fakePayrollDatabase) Clear() {
	db.employees = make(map[int]*Employee)
}
