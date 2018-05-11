package bridge

import (
	"errors"
	"strings"
	"testing"
)

func TestPrinterAPI1(t *testing.T) {
	api1 := PrinterImpl1{}

	err := api1.PrintMessage("Hello")

	if err != nil {
		t.Errorf("Error trying to use the API1 implementation: Message: %s\n", err.Error())
	}
}

type TestWriter struct {
	Msg string
}

func (t *TestWriter) Write(p []byte) (n int, err error) {
	n = len(p)

	if n > 0 {
		t.Msg = string(p)
		return n, nil
	}
	err = errors.New("Content received on Writer was empty")
	return
}

func TestPrinterAPI2WithNoIOWriter(t *testing.T) {
	api2 := PrinterImpl2{}

	err := api2.PrintMessage("Hello")

	if err != nil {
		expected := "You need to pass an io.Writer to PrinterImpl2"

		if !strings.Contains(err.Error(), expected) {
			t.Errorf("Error message was not correct.\ngot: %s\nexpected: %s\n", err.Error(), expected)
		}
	}
}

func TestPrinterAPI2WithIOWriter(t *testing.T) {
	testWriter := TestWriter{}
	api2 := PrinterImpl2{
		Writer: &testWriter,
	}
	expected := "Hello"
	err := api2.PrintMessage(expected)

	if err != nil {
		t.Errorf("Error trying to use the API2 implementation: %s\n", err.Error())
	}
	if testWriter.Msg != expected {
		t.Errorf("API2 did not write correctly to the io.Writer.\ngot: %s\nexpected: %s\n", testWriter.Msg, expected)
	}
}

func TestNormalPrinterWithPrinterImpl1(t *testing.T) {
	expected := "Hello io.Writer"

	normal := NormalPrinter{
		Msg:     expected,
		Printer: &PrinterImpl1{},
	}
	err := normal.Print()

	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestNormalPrinterWithPrinterImpl2(t *testing.T) {
	expected := "Hello io.Writer"
	testWriter := TestWriter{}
	normal := NormalPrinter{
		Msg: expected,
		Printer: &PrinterImpl2{
			Writer: &testWriter,
		},
	}
	err := normal.Print()

	if err != nil {
		t.Errorf(err.Error())
	}
	if testWriter.Msg != expected {
		t.Errorf("The expected message on the io.Writer doesn't match.\ngot: %s\nexpected: %s\n", testWriter.Msg, expected)
	}
}

func TestPacktPrinterWithPrinterImpl1(t *testing.T) {
	packt := PacktPrinter{
		Msg:     "",
		Printer: &PrinterImpl1{},
	}
	err := packt.Print()

	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestPacktPrinterWithPrinterImpl2(t *testing.T) {
	passed := "Hello io.Writer"
	expected := "Message from Packt: Hello io.Writer"
	testWriter := TestWriter{}
	packt := PacktPrinter{
		Msg: passed,
		Printer: &PrinterImpl2{
			Writer: &testWriter,
		},
	}
	err := packt.Print()

	if err != nil {
		t.Errorf(err.Error())
	}
	if testWriter.Msg != expected {
		t.Errorf("The expected message on the io.Writer doesn't match.\ngot: %s\nexpected: %s\n", testWriter.Msg, expected)
	}
}
