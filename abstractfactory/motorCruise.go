package abstractfactory

// CruiseMotor is a Motor with CruiseMotorType
type CruiseMotor struct{}

// GetType returns type of CruiseMotor
func (c *CruiseMotor) GetType() int {
	return CruiseMotorType
}

// NumWheels returns number of wheels of CruiseMotor
func (c *CruiseMotor) NumWheels() int {
	return 2
}

// NumSeats returns number of seats of CruiseMotor
func (c *CruiseMotor) NumSeats() int {
	return 2
}
