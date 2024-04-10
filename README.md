# Credit Card Validator

## Introduction

Implements the Luhn Algorithm

### Computing the check digit

If the number already contains the check digit, drop that digit to form the "payload". The check digit is most often the last digit.
With the payload, start from the rightmost digit. Moving left, double the value of every second digit (including the rightmost digit).
Sum the values of the resulting digits.
The check digit is calculated by (10 - (step3 mod 10)) mod 10. This is the smallest number (possibly zero) that must be added to step3 to make a multiple of 10.

### Validating the check digit

Drop the check digit (last digit) of the number to validate. (e.g. 17893729974 → 1789372997)
Calculate the check digit (see above)
Compare your result with the original check digit. If both numbers match, the result is valid. (e.g. (givenCheckDigit = calculatedCheckDigit) ⇔ (isValidCheckDigit)).
