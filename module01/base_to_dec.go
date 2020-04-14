package module01

import (
	"math"
	"strconv"
)

// BaseToDec takes in a number and the base it is currently
// in and returns the decimal equivalent as an integer.
//
// Eg:
//
//   BaseToDec("E", 16) => 14
//   BaseToDec("1110", 2) => 14
//
func BaseToDec(value string, base int) int {
	var result int
	for i, v := range Reversed(value) {
		intv, _ := strconv.Atoi(string(v))

		result = result + intv*int(math.Pow(float64(base), float64(i)))
	}
	return result
}

//Reversed function reverses a string
func Reversed(value string) string {
	var rev string
	for i := len(value) - 1; i >= 0; i-- {
		rev = rev + string(value[i])
	}
	return (rev)
}
