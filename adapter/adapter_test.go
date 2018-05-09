package adapter

import "testing"

func TestAdapter(t *testing.T) {
	msg := "Hello World!"
	adapter := PrinterAdapter{
		&OldPrinter{},
		msg,
	}
	expected := "Legacy Printer: Adapter: Hello World!"
	got := adapter.PrintStored()

	if got != expected {
		t.Errorf("Message didn't match, expected %s, but got %s", expected, got)
	}
}

func TestAdapterHasNoLegacyPrinter(t *testing.T) {
	msg := "Hello World!"
	adapter := PrinterAdapter{
		nil,
		msg,
	}
	expected := msg
	got := adapter.PrintStored()

	if got != expected {
		t.Errorf("Message didn't match, expected %s, but got %s", expected, got)
	}
}
