package payroll

import "payroll/domain"

type DirectMethod struct {
	Bank    string
	Account string
}

func (m *DirectMethod) Pay(pc *domain.Paycheck) {
}
