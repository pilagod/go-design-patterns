package abstractfactory

// SportMotor is a Motor with SportMotorType
type SportMotor struct{}

// GetType returns type of SportMotor
func (s *SportMotor) GetType() int {
	return SportMotorType
}

// NumWheels returns number of wheels of SportMotor
func (s *SportMotor) NumWheels() int {
	return 2
}

// NumSeats returns number of seats of SportMotor
func (s *SportMotor) NumSeats() int {
	return 1
}
