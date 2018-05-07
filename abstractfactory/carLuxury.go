package abstractfactory

// LuxuryCar is a Car with LuxuryCarType
type LuxuryCar struct{}

// GetType returns type of LuxuryCar
func (l *LuxuryCar) GetType() int {
	return LuxuryCarType
}

// NumDoors returns number of doors of LuxuryCar
func (l *LuxuryCar) NumDoors() int {
	return 4
}

// NumSeats returns number of seats of LuxuryCar
func (l *LuxuryCar) NumSeats() int {
	return 5
}

// NumWheels returns number of wheels of LuxuryCar
func (l *LuxuryCar) NumWheels() int {
	return 4
}
