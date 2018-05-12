package decorator

import (
	"fmt"
)

type PizzaMaker interface {
	Make() string
}

type NormalPizzaMaker struct{}

func (p *NormalPizzaMaker) Make() string {
	return "Pizza with the following ingredients:"
}

type OnionPizzaMaker struct {
	PizzaMaker PizzaMaker
}

func (o *OnionPizzaMaker) Make() string {
	return fmt.Sprintf("%s %s,", o.PizzaMaker.Make(), "onion")
}

type MeatPizzaMaker struct {
	PizzaMaker PizzaMaker
}

func (m *MeatPizzaMaker) Make() string {
	return fmt.Sprintf("%s %s,", m.PizzaMaker.Make(), "meat")
}
