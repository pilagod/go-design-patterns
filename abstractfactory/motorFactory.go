package abstractfactory

import (
	"fmt"
)

const (
	// SportMotorType is type of SportMotor
	SportMotorType = 1
	// CruiseMotorType is type of CruiseMotor
	CruiseMotorType = 2
)

// MotorFactory is a factory building Motor Vehicle
type MotorFactory struct{}

// Build builds a Motor Vehicle of given motor type
func (m *MotorFactory) Build(motorType int) (Vehicle, error) {
	switch motorType {
	case SportMotorType:
		return new(SportMotor), nil
	case CruiseMotorType:
		return new(CruiseMotor), nil
	default:
		return nil, fmt.Errorf("Motor of type %d not recognized", motorType)
	}
}

// GetType returns type of MotorFactory
func (m *MotorFactory) GetType() int {
	return MotorFactoryType
}
