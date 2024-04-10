package validator

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Validate(creditCardNumber string) bool {
	if !hasAllDigits(creditCardNumber) {
		return false
	}

	// check digit will be leftmost digit
	givenCheckDigit := string(creditCardNumber[len(creditCardNumber)-1])
	checkDigit := strconv.Itoa(computeCheckDigit(creditCardNumber[:len(creditCardNumber)-1]))

	return givenCheckDigit == checkDigit
}

func computeCheckDigit(accountNumber string) int {
	if !hasAllDigits(accountNumber) {
		return -1
	}

	payload := strings.Split(accountNumber, "")

	sum := 0
	startEven := len(payload) % 2
	for i := len(payload) - 1; i >= 0; i-- {
		if i%2 == startEven {
			digit, err := strconv.Atoi(payload[i])
			if err != nil {
				return -1
			}
			sum += (2 * digit)
		}
	}

	checkDigit := (10 - (sum % 10)) % 10

	return checkDigit
}

func ConstructCreditCardNumber(accountNumber string) string {
	if !hasAllDigits(accountNumber) {
		return ""
	}

	checkDigit := strconv.Itoa(computeCheckDigit(accountNumber))
	return fmt.Sprintf(`%s%s`, accountNumber, checkDigit)
}

func hasAllDigits(cardNumber string) bool {
	allNumbers := `^\d+$`
	re, err := regexp.Compile(allNumbers)
	if err != nil {
		fmt.Println("ERROR: ", err)
	}

	matched := re.MatchString(cardNumber)

	return matched
}
