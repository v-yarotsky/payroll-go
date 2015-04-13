package payroll

type DirectMethod struct {
	Bank    string
	Account string
}

func (m *DirectMethod) Pay(pc *Paycheck) {
}
