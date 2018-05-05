package builder

// MotorBuilder is a builder for Motor
type MotorBuilder struct {
	vp VehicleProduct
}

// SetWheels build 2 wheels for Motor
func (mb *MotorBuilder) SetWheels() VehicleBuilder {
	mb.vp.Wheels = 2
	return mb
}

// SetSeats build 2 seats for Motor
func (mb *MotorBuilder) SetSeats() VehicleBuilder {
	mb.vp.Seats = 2
	return mb
}

// SetStructure build a Motor structure for Motor
func (mb *MotorBuilder) SetStructure() VehicleBuilder {
	mb.vp.Structure = "Motor"
	return mb
}

// GetVehicle return the built Motor
func (mb *MotorBuilder) GetVehicle() VehicleProduct {
	return mb.vp
}
