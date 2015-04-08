package payroll

type CommissionedClassification struct {
	EmployeeID     int
	Salary         float64
	CommissionRate float64
}

func (c *CommissionedClassification) GetSalesReceipt(date int) (*SalesReceipt, error) {
	return GpayrollDatabase.GetSalesReceipt(c.EmployeeID, date)
}

func (c *CommissionedClassification) AddSalesReceipt(sr *SalesReceipt) error {
	return GpayrollDatabase.AddSalesReceipt(c.EmployeeID, sr)
}
