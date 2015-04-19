package payroll

import "payroll/domain"

type HoldMethod struct {
}

func (m *HoldMethod) Pay(pc *domain.Paycheck) {
	pc.SetField("Disposition", "Hold")
}
