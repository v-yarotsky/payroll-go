package main

import "fmt"
import . "payroll"

func main() {
	transactionSource := NewTextParserTransactionSource()
	for {
		t := transactionSource.GetTransaction()
		t.Execute()
		es := GpayrollDatabase.GetAllEmployees()
		for _, e := range es {
			fmt.Println(e)
		}
	}
}
