package payroll

import "fmt"
import "reflect"
import "time"
import "sort"

type HalpTransaction struct{}

func (h *HalpTransaction) Execute() error {
	fmt.Println("Available transactions:")
	txnNames := make([]string, 0, len(poop))
	for k, _ := range poop {
		txnNames = append(txnNames, k)
	}
	sort.Strings(txnNames)
	for _, s := range txnNames {
		fmt.Println(s)
	}
	return nil
}

var poop = map[string]func() Transaction{
	"add_commissioned_employee": func() Transaction {
		var d struct {
			EmployeeID     int     `ask:"Employee ID"`
			Name           string  `ask:"Employee Name"`
			Address        string  `ask:"Employee Address"`
			Salary         float64 `ask:"Salary"`
			CommissionRate float64 `ask:"Commission Rate"`
		}
		ask(&d)
		return NewAddCommissionedEmployeeTransaction(d.EmployeeID, d.Name, d.Address, d.Salary, d.CommissionRate)
	},
	"add_hourly_employee": func() Transaction {
		var d struct {
			EmployeeID int     `ask:"Employee ID"`
			Name       string  `ask:"Employee Name"`
			Address    string  `ask:"Employee Address"`
			HourlyRate float64 `ask:"Hourly Rate"`
		}
		ask(&d)
		return NewAddHourlyEmployeeTransaction(d.EmployeeID, d.Name, d.Address, d.HourlyRate)
	},
	"add_salaried_employee": func() Transaction {
		var d struct {
			EmployeeID int     `ask:"Employee ID"`
			Name       string  `ask:"Employee Name"`
			Address    string  `ask:"Employee Address"`
			Salary     float64 `ask:"Salary"`
		}
		ask(&d)
		return NewAddSalariedEmployeeTransaction(d.EmployeeID, d.Name, d.Address, d.Salary)
	},
	"change_address": func() Transaction {
		var d struct {
			EmployeeID int    `ask:"Employee ID"`
			Address    string `ask:"Employee Address"`
		}
		ask(&d)
		return NewChangeAddressTransaction(d.EmployeeID, d.Address)
	},
	"change_commissioned": func() Transaction {
		var d struct {
			EmployeeID     int     `ask:"Employee ID"`
			Salary         float64 `ask:"Salary"`
			CommissionRate float64 `ask:"Commission Rate"`
		}
		ask(&d)
		return NewChangeCommissionedTransaction(d.EmployeeID, d.Salary, d.CommissionRate)
	},
	"change_direct": func() Transaction {
		var d struct {
			EmployeeID int    `ask:"Employee ID"`
			Bank       string `ask:"Bank"`
			Account    string `ask:"Account"`
		}
		ask(&d)
		return NewChangeDirectTransaction(d.EmployeeID, d.Bank, d.Account)
	},
	"change_hold": func() Transaction {
		var d struct {
			EmployeeID int `ask:"Employee ID"`
		}
		ask(&d)
		return NewChangeHoldTransaction(d.EmployeeID)
	},
	"change_hourly": func() Transaction {
		var d struct {
			EmployeeID int     `ask:"Employee ID"`
			HourlyRate float64 `ask:"Hourly Rate"`
		}
		ask(&d)
		return NewChangeHourlyTransaction(d.EmployeeID, d.HourlyRate)
	},
	"change_mail": func() Transaction {
		var d struct {
			EmployeeID int    `ask:"Employee ID"`
			Address    string `ask:"Address"`
		}
		ask(&d)
		return NewChangeMailTransaction(d.EmployeeID, d.Address)
	},
	"change_member": func() Transaction {
		var d struct {
			EmployeeID int     `ask:"Employee ID"`
			MemberID   int     `ask:"Union Member ID"`
			Dues       float64 `ask:"Dues"`
		}
		ask(&d)
		return NewChangeMemberTransaction(d.EmployeeID, d.MemberID, d.Dues)
	},
	"change_name": func() Transaction {
		var d struct {
			EmployeeID int    `ask:"Employee ID"`
			Name       string `ask:"Employee Name"`
		}
		ask(&d)
		return NewChangeNameTransaction(d.EmployeeID, d.Name)
	},
	"change_salaried": func() Transaction {
		var d struct {
			EmployeeID int     `ask:"Employee ID"`
			Salary     float64 `ask:"Salary"`
		}
		ask(&d)
		return NewChangeSalariedTransaction(d.EmployeeID, d.Salary)
	},
	"change_unaffiliated": func() Transaction {
		var d struct {
			EmployeeID int `ask:"Employee ID"`
		}
		ask(&d)
		return NewChangeUnaffiliatedTransaction(d.EmployeeID)
	},
	"delete_employee": func() Transaction {
		var d struct {
			EmployeeID int `ask:"Employee ID"`
		}
		ask(&d)
		return NewDeleteEmployeeTransaction(d.EmployeeID)
	},
	"payday": func() Transaction {
		return NewPaydayTransaction(time.Now())
	},
	"sales_receipt": func() Transaction {
		var d struct {
			EmployeeID int       `ask:"Employee ID"`
			Date       time.Time `ask:"Date"`
			Amount     float64   `ask:"Amount"`
		}
		ask(&d)
		return NewSalesReceiptTransaction(d.EmployeeID, d.Date, d.Amount)
	},
	"service_charge": func() Transaction {
		var d struct {
			MemberID int       `ask:"Union Member ID"`
			Date     time.Time `ask:"Date"`
			Charge   float64   `ask:"Charge"`
		}
		ask(&d)
		return NewServiceChargeTransaction(d.MemberID, d.Date, d.Charge)
	},
	"time_card": func() Transaction {
		var d struct {
			EmployeeID int       `ask:"Employee ID"`
			Date       time.Time `ask:"Date"`
			Hours      float64   `ask:"Hours"`
		}
		ask(&d)
		return NewTimeCardTransaction(d.Date, d.Hours, d.EmployeeID)
	},
	"halp": func() Transaction {
		return &HalpTransaction{}
	},
}

type TransactionSource interface {
	GetTransaction() Transaction
}

type textParserTransactionSource struct {
}

func NewTextParserTransactionSource() *textParserTransactionSource {
	return &textParserTransactionSource{}
}

func (t *textParserTransactionSource) GetTransaction() Transaction {
	var tr Transaction

	for tr == nil {
		var tn struct {
			Name string `ask:"Transaction Name"`
		}
		ask(&tn)

		if trMaker, ok := poop[tn.Name]; ok {
			tr = trMaker()
		} else {
			fmt.Printf("Unknown transaction name \"%v\"\n", tn.Name)
			(&HalpTransaction{}).Execute()
		}
	}
	return tr
}

func ask(dataPtr interface{}) {
	s := reflect.ValueOf(dataPtr).Elem()
	typeOfS := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		ft := typeOfS.Field(i)
		question := ft.Tag.Get("ask")
		fmt.Printf(question + ": ")
		fmt.Scanln(f.Addr().Interface())
	}
}
