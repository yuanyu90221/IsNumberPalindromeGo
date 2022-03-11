package sol

import (
	"math"
)

func IsNumberPalindrome(input int) bool {
	quotient := input / 10
	count := 0
	for quotient > 0 {
		quotient /= 10
		count++
	}
	quotient = input
	reverseInput := 0
	digit := quotient
	for count >= 0 {
		digit = quotient % 10
		quotient /= 10
		reverseInput += int(math.Pow10(count)) * digit
		count--
	}
	return reverseInput == input
}
