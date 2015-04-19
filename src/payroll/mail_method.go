package payroll

import "payroll/domain"

type MailMethod struct {
	Address string
}

func (m *MailMethod) Pay(pc *domain.Paycheck) {
}
