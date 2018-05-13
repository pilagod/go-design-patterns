package main

import "fmt"

type One struct{}
type Two struct{}
type Three struct{}
type Four struct{}

func Sum(a, b interface{}) interface{} {
	switch a.(type) {
	case One:
		switch b.(type) {
		case One:
			return &Two{}
		case Two:
			return &Three{}
		case int:
			return b.(int) + 1
		default:
			return fmt.Errorf("Number not found")
		}
	case Two:
		switch b.(type) {
		case One:
			return &Three{}
		case Two:
			return &Four{}
		case int:
			return b.(int) + 2
		default:
			return fmt.Errorf("Number not found")
		}
	case int:
		switch b.(type) {
		case One:
		case Two:
			if a == 1 {
				return Sum(One{}, b)
			}
		case int:
			return a.(int) + b.(int)
		default:
			return fmt.Errorf("Number not found")
		}
	}
	return fmt.Errorf("Number not found")
}

func main() {
	fmt.Printf("%#v\n", Sum(One{}, Two{}))
	fmt.Printf("%#v\n", Sum(1, 2))
	fmt.Printf("%#v\n", Sum(One{}, 2))
}
