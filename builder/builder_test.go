package builder

import "testing"

func TestCarBuilder(t *testing.T) {
	md := new(ManufacturingDirector)
	cb := new(CarBuilder)
	md.SetBuilder(cb)
	md.Construct()

	car := cb.GetVehicle()

	if car.Wheels != 4 {
		t.Error("Car should have 4 wheels")
	}
	if car.Seats != 5 {
		t.Error("Car should have 5 seats")
	}
	if car.Structure != "Car" {
		t.Error("Car should have a structure of Car")
	}
}

func TestMotorBuilder(t *testing.T) {
	md := new(ManufacturingDirector)
	mb := new(MotorBuilder)
	md.SetBuilder(mb)
	md.Construct()

	motor := mb.GetVehicle()

	if motor.Wheels != 2 {
		t.Error("Motor should have 2 wheels")
	}
	if motor.Seats != 2 {
		t.Error("Motor should have 2 seats")
	}
	if motor.Structure != "Motor" {
		t.Error("Motor should have a structure of Motor")
	}
}
