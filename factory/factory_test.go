package factory

import (
	"strings"
	"testing"
)

func TestCashPayment(t *testing.T) {
	payment, err := GetPaymentMethod(Cash)

	if err != nil {
		t.Fatal("A payment method of type 'Cash' must exist")
	}
	msg := payment.Pay(10.30)

	if !strings.Contains(msg, "paid 10.30 in cash") {
		t.Error("Cash payment didn't pay properly")
	}
	t.Log("LOG:", msg)
}

func TestCreditCardPayment(t *testing.T) {
	payment, err := GetPaymentMethod(CreditCard)

	if err != nil {
		t.Fatal("A payment method of type 'CreditCard' must exist")
	}
	msg := payment.Pay(10.30)

	if !strings.Contains(msg, "paid 10.30 by credit card") {
		t.Error("Credit Card payment didn't pay properly")
	}
	t.Log("LOG:", msg)
}

func TestPaymentMethodNotExist(t *testing.T) {
	const SomeOtherPaymentMethod = -1

	_, err := GetPaymentMethod(SomeOtherPaymentMethod)

	if err == nil {
		t.Error("GetPaymentMethod should return error given unknown payment type")
	}
	t.Log("LOG:", err)
}
