package abstractfactory

import (
	"fmt"
)

const (
	// CarFactoryType is type of CarFactory
	CarFactoryType = 1
	// MotorFactoryType is type of MotorFactory
	MotorFactoryType = 2
)

// BuildVehicleFactory builds VehicleFactory of given factory type
func BuildVehicleFactory(ft int) (VehicleFactory, error) {
	switch ft {
	case CarFactoryType:
		return new(CarFactory), nil
	case MotorFactoryType:
		return new(MotorFactory), nil
	default:
		return nil, fmt.Errorf("Factory of type %d not recognized", ft)
	}
}
