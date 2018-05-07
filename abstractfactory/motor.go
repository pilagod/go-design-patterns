package abstractfactory

// Motor is an interface composing Vehicle
type Motor interface {
	Vehicle
	GetType() int
}
