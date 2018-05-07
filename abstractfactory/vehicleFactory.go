package abstractfactory

// Vehicle is an interface describing common vehicles
type Vehicle interface {
	NumSeats() int
	NumWheels() int
}

// VehicleFactory builds a Vehicle
type VehicleFactory interface {
	Build(vf int) (Vehicle, error)
	GetType() int
}
