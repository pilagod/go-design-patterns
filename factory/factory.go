package factory

import (
	"errors"
	"fmt"
)

const (
	// Cash payment type for Cash
	Cash = 1
	// CreditCard payment type for Credit Card
	CreditCard = 2
)

// GetPaymentMethod is a factory for PaymentMethod
func GetPaymentMethod(paymentType int) (PaymentMethod, error) {
	switch paymentType {
	case Cash:
		return &CashPaymentMethod{}, nil
	case CreditCard:
		return &CreditCardPaymentMethod{}, nil
	default:
		return nil, errors.New("Payment is not recognized")
	}
}

// PaymentMethod is an interface for paying
type PaymentMethod interface {
	Pay(amount float32) string
}

// CashPaymentMethod pays things in cash
type CashPaymentMethod struct{}

// Pay pays things in cash
func (c *CashPaymentMethod) Pay(amount float32) string {
	return fmt.Sprintf("You successfully paid %0.2f in cash", amount)
}

// CreditCardPaymentMethod pays things in credit card
type CreditCardPaymentMethod struct{}

// Pay pays things in credit card
func (c *CreditCardPaymentMethod) Pay(amount float32) string {
	return fmt.Sprintf("You successfully paid %0.2f by credit card", amount)
}
