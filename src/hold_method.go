package payroll

type HoldMethod struct {
}

func (m *HoldMethod) Pay(pc *Paycheck) {
	pc.SetField("Disposition", "Hold")
}
