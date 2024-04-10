package validator

import "testing"

func TestValidate(t *testing.T) {
	type test struct {
		cardNumber string
		want       bool
	}

	tests := []test{
		{"17893729974", true},
		{"17893729975", false},
		{"1789372994", false},
		{"abcd", false},
	}

	for _, tc := range tests {
		got := Validate(tc.cardNumber)
		if got != tc.want {
			t.Errorf(`Card Number: %s, want: %v, got: %v`, tc.cardNumber, tc.want, got)
		}
	}
}

func TestComputeCheckDigit(t *testing.T) {
	type test struct {
		cardNumber string
		want       int
	}

	tests := []test{
		{"1789372997", 4},
		{"178937299", 6},
		{"17893729", 2},
		{"1789372", 4},
		{"1789r372", -1},
	}

	for _, tc := range tests {
		got := computeCheckDigit(tc.cardNumber)
		if got != tc.want {
			t.Errorf(`Card Number: %s, want: %v, got: %v`, tc.cardNumber, tc.want, got)
		}
	}
}

func TestConstructCreditCardNumber(t *testing.T) {
	type test struct {
		accountNumber string
		want          string
	}

	tests := []test{
		{"1789372997", "17893729974"},
		{"178937299", "1789372996"},
		{"17893729", "178937292"},
		{"1789372", "17893724"},
		{"178sd9372", ""},
		{"abcdef", ""},
	}

	for _, tc := range tests {
		got := ConstructCreditCardNumber(tc.accountNumber)
		if got != tc.want {
			t.Errorf(`Card Number: %s, want: %v, got: %v`, tc.accountNumber, tc.want, got)
		}
	}
}
func TestHasAllDigits(t *testing.T) {
	type test struct {
		cardNumber string
		want       bool
	}

	tests := []test{
		{"1", true},
		{"1345346456", true},
		{"134534", true},
		{"134534 34534", false},
		{"11234f465", false},
		{"abcd", false},
	}

	for _, tc := range tests {
		got := hasAllDigits(tc.cardNumber)
		if got != tc.want {
			t.Errorf(`Card Number: %s, want: %v, got: %v`, tc.cardNumber, tc.want, got)
		}
	}
}

func BenchmarkValidate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Validate("17893729974")
	}
}

func BenchmarkComputeCheckDigit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		computeCheckDigit("1789372997")
	}
}

func BenchmarkConstructCreditCardNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConstructCreditCardNumber("1789372997")
	}
}

func BenchmarkHasAllDigits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		hasAllDigits("17893729974")
	}
}
