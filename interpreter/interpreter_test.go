package interpreter

import "testing"

func TestCalculateSimple(t *testing.T) {
	operation := "3 4 sum 2 sub"
	res, err := Calculate(operation)

	if err != nil {
		t.Error(err)
	}
	if res != 5 {
		t.Errorf("expected: %d, but got: %d", 5, res)
	}
}

func TestCalculateComplex(t *testing.T) {
	operation := "5 3 sub 8 mul 4 sum 5 div"
	res, err := Calculate(operation)

	if err != nil {
		t.Error(err)
	}
	if res != 4 {
		t.Errorf("expected: %d, but got: %d", 4, res)
	}
}

func TestInterpreterCalculateSimple(t *testing.T) {
	operation := "3 4 sum 2 sub"
	res, err := InterpreterCalculate(operation)

	if err != nil {
		t.Error(err)
	}
	if res != 5 {
		t.Errorf("expected: %d, but got: %d", 5, res)
	}
}

func TestInterpreterCalculateComplex(t *testing.T) {
	operation := "5 3 sub 8 mul 4 sum 5 div"
	res, err := InterpreterCalculate(operation)

	if err != nil {
		t.Error(err)
	}
	if res != 4 {
		t.Errorf("expected: %d, but got: %d", 4, res)
	}
}
