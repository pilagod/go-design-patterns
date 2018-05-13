package visitor

import "testing"

type TestHelper struct {
	Msg string
}

func (ts *TestHelper) Write(p []byte) (int, error) {
	ts.Msg = string(p)
	return len(p), nil
}

func TestMessageVisitor(t *testing.T) {
	testHelper := &TestHelper{}
	visitor := &MessageVisitor{}

	t.Run("MessageA", func(t *testing.T) {
		msg := MessageA{
			Msg:    "Hello World",
			Output: testHelper,
		}
		msg.Accept(visitor)
		msg.Print()
		expected := "A: Hello World (Visited A)"

		if testHelper.Msg != expected {
			t.Errorf("expected: %s, but got: %s", expected, testHelper.Msg)
		}
	})

	t.Run("MessageB", func(t *testing.T) {
		msg := MessageB{
			Msg:    "Hello World",
			Output: testHelper,
		}
		msg.Accept(visitor)
		msg.Print()
		expected := "B: Hello World (Visited B)"

		if testHelper.Msg != expected {
			t.Errorf("expected: %s, but got: %s", expected, testHelper.Msg)
		}
	})
}

func TestProductVisitor(t *testing.T) {
	products := make([]ProductVisitable, 2)
	products[0] = &Rice{
		Product: Product{
			Price: 32.0,
			Name:  "Some rice",
		},
	}
	products[1] = &Pasta{
		Product: Product{
			Price: 40.0,
			Name:  "Some pasta",
		},
	}
	priceVisitor := &PriceVisitor{}

	for _, product := range products {
		product.Accept(priceVisitor)
	}
	if priceVisitor.Sum != 72.0 {
		t.Errorf("expected: %.2f, but got: %.2f", 72.0, priceVisitor.Sum)
	}
	nameVisitor := &ProductListVisitor{}

	for _, product := range products {
		product.Accept(nameVisitor)
	}
	if nameVisitor.ProductList != "Some rice, Some pasta" {
		t.Errorf("expected: %s, but got: %s", "Some rice, Some pasta", nameVisitor.ProductList)
	}
}
