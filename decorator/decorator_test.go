package decorator

import (
	"strings"
	"testing"
)

func TestNormalPizzaMaker(t *testing.T) {
	pizzaMaker := &NormalPizzaMaker{}
	pizza := pizzaMaker.Make()
	expected := "Pizza with the following ingredients:"

	if !strings.Contains(pizza, expected) {
		t.Errorf("When calling the add ingredient of the pizza decorator it must return text\ngot: %s\nexpected: %s", pizza, expected)
	}
}

func TestOnionPizzaMaker(t *testing.T) {
	onionPizzaMaker := &OnionPizzaMaker{&NormalPizzaMaker{}}
	onionPizza := onionPizzaMaker.Make()

	if !strings.Contains(onionPizza, "onion") {
		t.Errorf("OnionPizzaMaker must add onion to pizza, but got: %s", onionPizza)
	}
}

func TestMeatPizzaMaker(t *testing.T) {
	meatPizzaMaker := &MeatPizzaMaker{&NormalPizzaMaker{}}
	meatPizza := meatPizzaMaker.Make()

	if !strings.Contains(meatPizza, "meat") {
		t.Errorf("MeatPizzaMaker must add meat to pizza, but got %s", meatPizza)
	}
}

func TestOnionMeatPizzaMaker(t *testing.T) {
	pizzaMaker := &OnionPizzaMaker{
		&MeatPizzaMaker{
			&NormalPizzaMaker{},
		},
	}
	pizza := pizzaMaker.Make()
	expected := "Pizza with the following ingredients: meat, onion"

	if !strings.Contains(pizza, expected) {
		t.Errorf("OnionMeatPizzaMaker must add meat and onion to pizza, but got: %s", pizza)
	}
}
