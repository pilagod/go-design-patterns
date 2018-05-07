package prototype

import "testing"

func TestGetShirtClonerCloneShirtProperly(t *testing.T) {
	testCases := []struct {
		Color ShirtColor
		Proto *Shirt
	}{
		{White, whitePrototype},
		{Black, blackPrototype},
		{Blue, bluePrototype},
	}

	for _, testCase := range testCases {
		t.Log("---------------------------------------------")
		testGetShirtClonerCloneSpecificShirtProperly(t, testCase.Color, testCase.Proto)
	}
	t.Log("---------------------------------------------")
}

func testGetShirtClonerCloneSpecificShirtProperly(t *testing.T, color ShirtColor, proto *Shirt) {
	shirtCache := GetShirtCloner()

	if shirtCache == nil {
		t.Fatal("Received cache was nil")
	}

	// item1

	item1, err := shirtCache.GetClone(color)

	if err != nil {
		t.Fatal(err)
	}
	if item1 == proto {
		t.Fatalf("item1 cannot be equal to the %b prototype", color)
	}
	shirt1, ok := item1.(*Shirt)

	if !ok {
		t.Fatal("Type assertion for shirt1 couldn't be done successfully")
	}

	if *shirt1 != *proto {
		t.Errorf("shirt should clone from proto %v", *proto)
	}

	// item2

	item2, err := shirtCache.GetClone(color)

	if err != nil {
		t.Fatal(err)
	}
	shirt2, ok := item2.(*Shirt)

	if !ok {
		t.Fatal("Type assertion for shirt2 couldn't be done successfully")
	}

	// update shirt1 SKU
	shirt1.SKU = "abbcc"

	if shirt1.SKU == shirt2.SKU {
		t.Error("SKU of shirt1 should be different from shirt2's")
	}
	t.Logf("LOG: %s", shirt1.GetInfo())
	t.Logf("LOG: %s", shirt2.GetInfo())
	t.Logf("LOG: The memory positions of the shirts should be different %p != %p", &shirt1, &shirt2)
}
