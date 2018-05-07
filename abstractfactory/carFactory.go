package abstractfactory

import (
	"fmt"
)

const (
	// LuxuryCarType is type of LuxuryCar
	LuxuryCarType = 1
	// FamilyCarType is type of FamilyCar
	FamilyCarType = 2
)

// CarFactory is a factory building Car Vehicle
type CarFactory struct{}

// Build builds a Car Vehicle of given car type
func (c *CarFactory) Build(carType int) (Vehicle, error) {
	switch carType {
	case LuxuryCarType:
		return new(LuxuryCar), nil
	case FamilyCarType:
		return new(FamilyCar), nil
	default:
		return nil, fmt.Errorf("Car of type %d not recognized", carType)
	}
}

// GetType returns type of CarFactory
func (c *CarFactory) GetType() int {
	return CarFactoryType
}
