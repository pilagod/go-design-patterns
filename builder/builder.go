package builder

// VehicleBuilder is an interface for building VehicleProduct
type VehicleBuilder interface {
	SetWheels() VehicleBuilder
	SetSeats() VehicleBuilder
	SetStructure() VehicleBuilder
	GetVehicle() VehicleProduct
}

// VehicleProduct is a struct for real world vehicle
type VehicleProduct struct {
	Wheels    int
	Seats     int
	Structure string
}

// ManufacturingDirector directs building process of vehicle product
type ManufacturingDirector struct {
	builder VehicleBuilder
}

// Construct build a vehicle based on builder on ManufacturingDirector
func (md *ManufacturingDirector) Construct() {
	md.builder.SetWheels().SetSeats().SetStructure()
}

// SetBuilder set a desired vehicle builder on ManufacturingDirector
func (md *ManufacturingDirector) SetBuilder(vb VehicleBuilder) {
	md.builder = vb
}
