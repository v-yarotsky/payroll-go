package payroll

type changeMailTransaction struct {
	BaseChangeMethodTransaction
	Address string
}

func NewChangeMailTransaction(empId int, address string) *changeMailTransaction {
	return &changeMailTransaction{BaseChangeMethodTransaction{empId}, address}
}

func (t *changeMailTransaction) GetMethod() PaymentMethod {
	return &MailMethod{t.Address}
}

func (t *changeMailTransaction) Execute() error {
	return t.BaseChangeMethodTransaction.Execute(t)
}
