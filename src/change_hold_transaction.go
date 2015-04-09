package payroll

type changeHoldTransaction struct {
	BaseChangeMethodTransaction
}

func NewChangeHoldTransaction(empId int) *changeHoldTransaction {
	return &changeHoldTransaction{BaseChangeMethodTransaction{empId}}
}

func (t *changeHoldTransaction) GetMethod() PaymentMethod {
	return &HoldMethod{}
}

func (t *changeHoldTransaction) Execute() error {
	return t.BaseChangeMethodTransaction.Execute(t)
}
