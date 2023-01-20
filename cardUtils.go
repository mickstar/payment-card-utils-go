package CardUtils

import (
	"fmt"
	"github.com/mickstar/payment-card-utils-go/Scheme"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// maskPan Returns a masked version of the card number
// this is the first 6 digits and the last 4 digits,
// with the rest replaced by asterisks.
// if the pan is less than 10 digits, the PAN is returned without modification.
func MaskPan(pan string) string {
	return MaskPanWithCharacter(pan, '*')
}

// same as maskPan except using the utf8 character to mask the pan
func MaskPanWithCharacter(pan string, maskCharacter uint8) string {
	// mask pan
	if len(pan) <= 10 {
		return pan
	}

	return pan[0:6] + strings.Repeat(string(maskCharacter), len(pan)-10) + pan[len(pan)-4:]
}

func GetCardScheme(pan string) Scheme.Scheme {
	// test for visa
	if strings.HasPrefix(pan, "4") {
		return Scheme.Visa
	}
	// test for Amex
	if strings.HasPrefix(pan, "34") || strings.HasPrefix(pan, "37") {
		return Scheme.AmericanExpress
	}
	// test for MasterCard (51 - 55)
	if strings.HasPrefix(pan, "51") ||
		strings.HasPrefix(pan, "52") ||
		strings.HasPrefix(pan, "53") ||
		strings.HasPrefix(pan, "54") ||
		strings.HasPrefix(pan, "55") {
		return Scheme.MasterCard
	}

	// test for Diners
	if strings.HasPrefix(pan, "36") || strings.HasPrefix(pan, "38") {
		return Scheme.DinersClub
	}

	// test for JCB
	if strings.HasPrefix(pan, "35") {
		return Scheme.JCB
	}

	// test for Discover
	if strings.HasPrefix(pan, "6011") || strings.HasPrefix(pan, "65") {
		return Scheme.Discover
	}

	//test for BP Card
	if strings.HasPrefix(pan, "7052") || strings.HasPrefix(pan, "7050") {
		return Scheme.BPCard
	}

	//test for union pay 6200**-6250**
	if strings.HasPrefix(pan, "62") && pan[2] >= '0' && pan[2] <= '5' &&
		!(pan[2] == '5' && pan[3] != '0') {
		return Scheme.UnionPay
	}

	return Scheme.Unknown
}

// GenerateRandomPanOfLength generates a random pan of the specified length
// This PAN is guaranteed to pass the Luhn check.
func GenerateRandomPanOfLength(length int) string {
	// seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// generate 9 random numbers
	// each one within the range 0-9
	nums := make([]int, length-1)
	for i := 0; i < 9; i++ {
		nums[i] = rand.Intn(10)
	}
	// create the PAN string
	pan := ""
	for _, n := range nums {
		pan += fmt.Sprintf("%d", n)
	}

	for i := 0; i < 10; i++ {
		if LuhnCheck(pan + strconv.Itoa(i)) {
			return pan + strconv.Itoa(i)
		}
	}
	// should never happen.
	panic("Could not generate a valid PAN from base " + pan)
}

// GenerateRandomPanOfLength generates a random pan of 16 digits
// This PAN is guaranteed to pass the Luhn check.
func GenerateRandomPan() string {
	return GenerateRandomPanOfLength(16)
}

// LuhnCheck performs a Luhn check on the card number
func LuhnCheck(pan string) bool {
	// check every digit is int
	for _, c := range pan {
		if c < '0' || c > '9' {
			return false
		}
	}

	controlDigit := pan[len(pan)-1]
	sum := 0
	shouldDouble := true
	for i := len(pan) - 2; i >= 0; i-- {
		digit := pan[i] - '0'
		if shouldDouble {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += int(digit)
		shouldDouble = !shouldDouble
	}
	return (10-sum%10)%10 == int(controlDigit-'0')
}

// ValidityCheck verifies that the PAN matches the Scheme's card length,
// that the PAN is all numbers
// and that the PAN passes the Luhn check.
func ValidityCheck(pan string) bool {
	// check if pan is valid
	if len(pan) < 14 || len(pan) > 19 {
		return false
	}

	if !Scheme.LengthCheckForScheme(GetCardScheme(pan), len(pan)) {
		return false
	}

	return LuhnCheck(pan)
}
