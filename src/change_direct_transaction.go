package payroll

type changeDirectTransaction struct {
	BaseChangeMethodTransaction
	Bank    string
	Account string
}

func NewChangeDirectTransaction(empId int, bank string, account string) *changeDirectTransaction {
	return &changeDirectTransaction{BaseChangeMethodTransaction{empId}, bank, account}
}

func (t *changeDirectTransaction) GetMethod() PaymentMethod {
	return &DirectMethod{t.Bank, t.Account}
}

func (t *changeDirectTransaction) Execute() error {
	return t.BaseChangeMethodTransaction.Execute(t)
}
