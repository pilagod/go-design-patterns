package abstractfactory

// Car is an interface composing Vehicle
type Car interface {
	Vehicle
	GetType() int
	NumDoors() int
}
