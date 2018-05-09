package adapter

import (
	"fmt"
)

// LegacyPrinter is an interface stands for old printers
type LegacyPrinter interface {
	Print(s string) string
}

// OldPrinter implements LegacyPrinter
type OldPrinter struct{}

// Print prints a new string from given string
func (o *OldPrinter) Print(s string) string {
	return fmt.Sprintf("Legacy Printer: %s", s)
}

// ModernPrinter is a interface stands for new printers
type ModernPrinter interface {
	PrintStored() string
}

// PrinterAdapter implements ModernPrinter, adapts LegacyPrinter and ModernPrinter
type PrinterAdapter struct {
	LegacyPrinter LegacyPrinter
	Msg           string
}

// PrintStored uses ModernPrinter's signature to print stored string
func (pa *PrinterAdapter) PrintStored() string {
	if pa.LegacyPrinter == nil {
		return pa.Msg
	}
	msg := fmt.Sprintf("Adapter: %s", pa.Msg)
	return pa.LegacyPrinter.Print(msg)
}
