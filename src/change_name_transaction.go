package payroll

type changeNameTransaction struct {
	BaseChangeEmployeeTransaction
	NewName string
}

func NewChangeNameTransaction(empId int, newName string) *changeNameTransaction {
	return &changeNameTransaction{BaseChangeEmployeeTransaction{empId}, newName}
}

func (t *changeNameTransaction) Execute() error {
	return t.BaseChangeEmployeeTransaction.Execute(t)
}

func (t *changeNameTransaction) Change(e *Employee) error {
	e.Name = t.NewName
	return nil
}
