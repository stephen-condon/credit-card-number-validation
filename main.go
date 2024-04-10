package main

import (
	"creditcardvalidator/validator"
	"fmt"
)

func main() {
	accountNumber := "1789372997"
	newCreditCardNumber := validator.ConstructCreditCardNumber(accountNumber)
	fmt.Printf("Account Number: %s; CreditCardNumber: %s", accountNumber, newCreditCardNumber)

	fmt.Println()
	fmt.Println()
	fmt.Println("Validate Credit Card Number:")
	isValid := validator.Validate(newCreditCardNumber)
	if isValid {
		fmt.Println("Credit Card Number is valid!")
	} else {
		fmt.Println("Credit Card Number is invalid!")
	}
}
