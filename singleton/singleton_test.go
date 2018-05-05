package singleton

import (
	"os"
	"testing"
)

// https://stackoverflow.com/questions/23729790/how-can-i-do-test-setup-using-the-testing-package-in-go?utm_medium=organic&utm_source=google_rich_qa&utm_campaign=google_rich_qa
func TestMain(m *testing.M) {
	println("before")
	retCode := m.Run()
	println("after")
	os.Exit(retCode)
}

func TestGetInstanceReturnsDefaultCounterWhenFirstCalled(t *testing.T) {
	counter := GetInstance()

	if counter.GetCurrentCount() != 0 {
		t.Error("GetInstance should return a Counter with property count initialized to 0 when first called")
	}
}

func TestGetInstanceAlwaysReturnsSameCounter(t *testing.T) {
	counter1 := GetInstance()
	counter2 := GetInstance()

	if counter1 != counter2 {
		t.Error("GetInstance should always return a same Counter")
	}
}

func TestIncreaseAddToCountInCounterByOne(t *testing.T) {
	counter := GetInstance()

	counter.Increase()

	if counter.GetCurrentCount() != 1 {
		t.Error("Increase should add 1 to property count in Counter")
	}
}
