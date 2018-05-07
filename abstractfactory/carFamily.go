package abstractfactory

// FamilyCar is a Car with FamilyCarType
type FamilyCar struct{}

// GetType returns type of FamilyCar
func (f *FamilyCar) GetType() int {
	return FamilyCarType
}

// NumDoors returns number of doors of FamilyCar
func (f *FamilyCar) NumDoors() int {
	return 5
}

// NumSeats returns number of seats of FamilyCar
func (f *FamilyCar) NumSeats() int {
	return 5
}

// NumWheels returns number of wheels of FamilyCar
func (f *FamilyCar) NumWheels() int {
	return 4
}
