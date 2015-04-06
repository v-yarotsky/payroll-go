package payroll

type HourlyClassification struct {
	EmployeeID int
	HourlyRate float64
}

func (c *HourlyClassification) GetTimeCard(date int) (*TimeCard, error) {
	return GpayrollDatabase.GetTimeCard(c.EmployeeID, date)
}

func (c *HourlyClassification) AddTimeCard(tc *TimeCard) {
	GpayrollDatabase.AddTimeCard(c.EmployeeID, tc)
}
