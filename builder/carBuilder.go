package builder

// CarBuilder is a builder for Car
type CarBuilder struct {
	vp VehicleProduct
}

// SetWheels build 4 wheels for Car
func (cb *CarBuilder) SetWheels() VehicleBuilder {
	cb.vp.Wheels = 4
	return cb
}

// SetSeats build 5 seats for Car
func (cb *CarBuilder) SetSeats() VehicleBuilder {
	cb.vp.Seats = 5
	return cb
}

// SetStructure build a Car structure for Car
func (cb *CarBuilder) SetStructure() VehicleBuilder {
	cb.vp.Structure = "Car"
	return cb
}

// GetVehicle return a built Car
func (cb *CarBuilder) GetVehicle() VehicleProduct {
	return cb.vp
}
