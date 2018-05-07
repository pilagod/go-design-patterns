package abstractfactory

import "testing"

func TestBuildVehicleFactoryShouldReturnCarFactoryGivenCarFactoryType(t *testing.T) {
	carFactory, err := BuildVehicleFactory(CarFactoryType)

	if err != nil {
		t.Fatal(err)
	}
	if carFactory.GetType() != CarFactoryType {
		t.Error("BuildVehicleFactory should return a CarFactory given factory type of CarFactoryType")
	}
}

func TestBuildVehicleFactoryShouldReturnCarFactoryThatBuildsLuxuryCar(t *testing.T) {
	carFactory, _ := BuildVehicleFactory(CarFactoryType)
	carVehicle, err := carFactory.Build(LuxuryCarType)

	if err != nil {
		t.Fatal(err)
	}
	luxuryCar, ok := carVehicle.(Car)

	if !ok {
		t.Fatal("CarFactory should return a Car Vehicle")
	}
	if luxuryCar.GetType() != LuxuryCarType {
		t.Fatal("CarFactory should return a LuxuryCar given LuxuryCarType")
	}
	if luxuryCar.NumDoors() != 4 {
		t.Error("LuxuryCar should have 4 doors")
	}
	if luxuryCar.NumSeats() != 5 {
		t.Error("LuxuryCar should have 5 seats")
	}
	if luxuryCar.NumWheels() != 4 {
		t.Error("LuxuryCar should have 4 wheels")
	}
}

func TestBuildVehicleFactoryShouldReturnCarFactoryThatBuildsFamilyCar(t *testing.T) {
	carFactory, _ := BuildVehicleFactory(CarFactoryType)
	carVehicle, err := carFactory.Build(FamilyCarType)

	if err != nil {
		t.Fatal(err)
	}
	familyCar, ok := carVehicle.(Car)

	if !ok {
		t.Fatal("CarFactory should return a Car Vehicle")
	}
	if familyCar.GetType() != FamilyCarType {
		t.Fatal("CarFactory should return a FamilyCar given FamilyCarType")
	}
	if familyCar.NumDoors() != 5 {
		t.Error("FamilyCar should have 5 doors")
	}
	if familyCar.NumSeats() != 5 {
		t.Error("FamilyCar should have 5 seats")
	}
	if familyCar.NumWheels() != 4 {
		t.Error("FamilyCar should have 4 wheels")
	}
}

func TestBuildVehicleFactoryShouldReturnMotorFactoryGivenCarFactoryType(t *testing.T) {
	motorFactory, err := BuildVehicleFactory(MotorFactoryType)

	if err != nil {
		t.Fatal(err)
	}
	if motorFactory.GetType() != MotorFactoryType {
		t.Error("BuildVehicleFactory should return a MotorFactory given factory type of MotorFactoryType")
	}
}

func TestBuildVehicleFactoryShouldReturnMotorFactoryThatBuildsSportMotor(t *testing.T) {
	motorFactory, _ := BuildVehicleFactory(MotorFactoryType)
	motorVehicle, err := motorFactory.Build(SportMotorType)

	if err != nil {
		t.Fatal(err)
	}
	sportMotor, ok := motorVehicle.(Motor)

	if ok == false {
		t.Fatal("MotorFactory should return a Motor Vehicle")
	}
	if sportMotor.GetType() != SportMotorType {
		t.Fatal("MotorFactory should return a SportMotor given SportMotorType")
	}
	if sportMotor.NumSeats() != 1 {
		t.Error("SportMotor should have 1 seats")
	}
	if sportMotor.NumWheels() != 2 {
		t.Error("SportMotor should have 2 wheels")
	}
}

func TestBuildVehicleFactoryShouldReturnMotorFactoryThatBuildsCruiseMotor(t *testing.T) {
	motorFactory, _ := BuildVehicleFactory(MotorFactoryType)
	motorVehicle, err := motorFactory.Build(CruiseMotorType)

	if err != nil {
		t.Fatal(err)
	}
	cruiseMotor, ok := motorVehicle.(Motor)

	if ok == false {
		t.Fatal("MotorFactory should return a Motor Vehicle")
	}
	if cruiseMotor.GetType() != CruiseMotorType {
		t.Fatal("MotorFactory should return a CruiseMotor given CruiseMotorType")
	}
	if cruiseMotor.NumSeats() != 2 {
		t.Error("SportMotor should have 2 seats")
	}
	if cruiseMotor.NumWheels() != 2 {
		t.Error("SportMotor should have 2 wheels")
	}
}
