package payroll

import "testing"
import "time"

func TestPaySingleSalariedEmployee(t *testing.T) {
	empId := 1
	addTr := NewAddSalariedEmployeeTransaction(empId, "Bill", "Home", 1000.0)
	addTr.Execute()

	payDate := parseDate("2001-Nov-30")
	tr := NewPaydayTransaction(payDate)
	tr.Execute()

	pc, err := tr.GetPaycheck(empId)
	if err != nil {
		t.Fatalf("salaried employee should have had a paycheck at the end of the month!")
	}

	validatePaycheck(pc, payDate, 1000.0, 0.0, 1000.0, "Hold", t)
}

func TestPaySingleSalariedEmployeeOnWrongDate(t *testing.T) {
	empId := 1
	addTr := NewAddSalariedEmployeeTransaction(empId, "Bill", "Home", 1000.0)
	addTr.Execute()

	payDate := parseDate("2001-Nov-15")
	tr := NewPaydayTransaction(payDate)
	tr.Execute()

	pc, err := tr.GetPaycheck(empId)
	if pc != nil || err == nil {
		t.Fatalf("expected employee to get no paycheck, got %v (error: %v)", pc, err)
	}
}

func TestPaySingleHourlyEmployeeNoTimeCards(t *testing.T) {
	empId := 2
	addTr := NewAddHourlyEmployeeTransaction(empId, "Bill", "Home", 10.0)
	addTr.Execute()

	payDate := parseDate("2001-Nov-09")
	tr := NewPaydayTransaction(payDate)
	tr.Execute()

	pc, err := tr.GetPaycheck(empId)
	if err != nil {
		t.Fatalf("hourly employee should have had a paycheck on Friday!")
	}

	validatePaycheck(pc, payDate, 0.0, 0.0, 0.0, "Hold", t)
}

func TestPaySingleHourlyEmployeeOneTimeCard(t *testing.T) {
	empId := 2
	addTr := NewAddHourlyEmployeeTransaction(empId, "Bill", "Home", 15.25)
	addTr.Execute()

	timeCardTr := NewTimeCardTransaction(parseDate("2001-Nov-08"), 2.0, empId)
	timeCardTr.Execute()

	payDate := parseDate("2001-Nov-09")
	tr := NewPaydayTransaction(payDate)
	tr.Execute()

	pc, err := tr.GetPaycheck(empId)
	if err != nil {
		t.Fatalf("hourly employee should have had a paycheck on Friday!")
	}

	validatePaycheck(pc, payDate, 30.5, 0.0, 30.5, "Hold", t)
}

func TestPaySingleHourlyEmployeeOvertimeOneTimeCard(t *testing.T) {
	empId := 2
	addTr := NewAddHourlyEmployeeTransaction(empId, "Bill", "Home", 15.25)
	addTr.Execute()

	timeCardTr := NewTimeCardTransaction(parseDate("2001-Nov-08"), 9.0, empId)
	timeCardTr.Execute()

	payDate := parseDate("2001-Nov-09")
	tr := NewPaydayTransaction(payDate)
	tr.Execute()

	pc, err := tr.GetPaycheck(empId)
	if err != nil {
		t.Fatalf("hourly employee should have had a paycheck on Friday!")
	}

	pay := 8*15.25 + 1*1.5*15.25
	validatePaycheck(pc, payDate, pay, 0.0, pay, "Hold", t)
}

func TestPaySingleHourlyEmployeeOnWrongDate(t *testing.T) {
	empId := 2
	addTr := NewAddHourlyEmployeeTransaction(empId, "Bill", "Home", 15.25)
	addTr.Execute()

	timeCardTr1 := NewTimeCardTransaction(parseDate("2001-Nov-07"), 2.0, empId)
	timeCardTr1.Execute()

	timeCardTr2 := NewTimeCardTransaction(parseDate("2001-Nov-09"), 5.0, empId)
	timeCardTr2.Execute()

	payDate := parseDate("2001-Nov-09")
	tr := NewPaydayTransaction(payDate)
	tr.Execute()

	pc, err := tr.GetPaycheck(empId)
	if err != nil {
		t.Fatalf("hourly employee should have had a paycheck on Friday!")
	}

	pay := 2*15.25 + 5*15.25
	validatePaycheck(pc, payDate, pay, 0.0, pay, "Hold", t)
}

func TestPaySingleHourlyEmployeeWithTimeCardsSpanningTwoPeriods(t *testing.T) {
	empId := 2
	addTr := NewAddHourlyEmployeeTransaction(empId, "Bill", "Home", 15.25)
	addTr.Execute()

	timeCardTr1 := NewTimeCardTransaction(parseDate("2001-Nov-01"), 2.0, empId) // previous pay period
	timeCardTr1.Execute()

	timeCardTr2 := NewTimeCardTransaction(parseDate("2001-Nov-09"), 5.0, empId)
	timeCardTr2.Execute()

	payDate := parseDate("2001-Nov-09")
	tr := NewPaydayTransaction(payDate)
	tr.Execute()

	pc, err := tr.GetPaycheck(empId)
	if err != nil {
		t.Fatalf("hourly employee should have had a paycheck on Friday!")
	}

	pay := 5 * 15.25
	validatePaycheck(pc, payDate, pay, 0.0, pay, "Hold", t)
}

func TestPaySingleCommissionedEmployeeWithNoSalesReceipts(t *testing.T) {
	empId := 3
	addTr := NewAddCommissionedEmployeeTransaction(empId, "Bill", "Home", 1000.0, 0.05)
	addTr.Execute()

	payDate := parseDate("2001-Nov-15")
	tr := NewPaydayTransaction(payDate)
	tr.Execute()

	pc, err := tr.GetPaycheck(empId)
	if err != nil {
		t.Fatalf("commissioned employee should have had a paycheck on second Friday!")
	}

	validatePaycheck(pc, payDate, 1000.0, 0.0, 1000.0, "Hold", t)
}

func TestPaySingleCommissionedEmployeeWithOneSalesReceipt(t *testing.T) {
	empId := 3
	addTr := NewAddCommissionedEmployeeTransaction(empId, "Bill", "Home", 1000.0, 0.05)
	addTr.Execute()

	salesReceiptTr := NewSalesReceiptTransaction(empId, parseDate("2001-Nov-14"), 10000.0)
	salesReceiptTr.Execute()

	payDate := parseDate("2001-Nov-15")
	tr := NewPaydayTransaction(payDate)
	tr.Execute()

	pc, err := tr.GetPaycheck(empId)
	if err != nil {
		t.Fatalf("commissioned employee should have had a paycheck on second Friday!")
	}

	pay := 1000.0 + 10000.0*0.05
	validatePaycheck(pc, payDate, pay, 0.0, pay, "Hold", t)
}

func TestPaySingleCommissionedEmployeeWithTwoSalesReceipts(t *testing.T) {
	empId := 3
	addTr := NewAddCommissionedEmployeeTransaction(empId, "Bill", "Home", 1000.0, 0.05)
	addTr.Execute()

	salesReceiptTr1 := NewSalesReceiptTransaction(empId, parseDate("2001-Nov-12"), 5000.0)
	salesReceiptTr1.Execute()

	salesReceiptTr2 := NewSalesReceiptTransaction(empId, parseDate("2001-Nov-14"), 10000.0)
	salesReceiptTr2.Execute()

	payDate := parseDate("2001-Nov-15")
	tr := NewPaydayTransaction(payDate)
	tr.Execute()

	pc, err := tr.GetPaycheck(empId)
	if err != nil {
		t.Fatalf("commissioned employee should have had a paycheck on second Friday!")
	}

	pay := 1000.0 + (5000.0+10000.0)*0.05
	validatePaycheck(pc, payDate, pay, 0.0, pay, "Hold", t)
}

func TestPaySingleCommissionedEmployeeWithSalesReceiptsSpanningTwoPeriods(t *testing.T) {
	empId := 3
	addTr := NewAddCommissionedEmployeeTransaction(empId, "Bill", "Home", 1000.0, 0.05)
	addTr.Execute()

	salesReceiptTr1 := NewSalesReceiptTransaction(empId, parseDate("2001-Oct-12"), 5000.0) // Previous pay period
	salesReceiptTr1.Execute()

	salesReceiptTr2 := NewSalesReceiptTransaction(empId, parseDate("2001-Nov-14"), 10000.0)
	salesReceiptTr2.Execute()

	payDate := parseDate("2001-Nov-15")
	tr := NewPaydayTransaction(payDate)
	tr.Execute()

	pc, err := tr.GetPaycheck(empId)
	if err != nil {
		t.Fatalf("commissioned employee should have had a paycheck on second Friday!")
	}

	pay := 1000.0 + 10000.0*0.05
	validatePaycheck(pc, payDate, pay, 0.0, pay, "Hold", t)
}

func validatePaycheck(pc *Paycheck, payDate time.Time, grossPay, deductions, netPay float64, disposition string, t *testing.T) {
	if pc.PayDate != payDate {
		t.Fatalf("expected paycheck to be dated %v, was dated %v", payDate, pc.PayDate)
	}

	if !floatEquals(pc.GrossPay, grossPay) {
		t.Fatalf("expected paycheck gross pay to eq %v, got %v", grossPay, pc.GrossPay)
	}

	if v, _ := pc.GetField("Disposition"); v != disposition {
		t.Fatalf("expected paycheck disposition to be %v", disposition)
	}

	if !floatEquals(pc.Deductions, deductions) {
		t.Fatalf("expected paycheck to have %v in deductions, got %v", deductions, pc.Deductions)
	}

	if !floatEquals(pc.NetPay, netPay) {
		t.Fatalf("expected paycheck net pay to eq %v, got %v", netPay, pc.NetPay)
	}
}
