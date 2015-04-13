package payroll

type MailMethod struct {
	Address string
}

func (m *MailMethod) Pay(pc *Paycheck) {
}
