package utility

import (
	"math"
	"strconv"
)

const expresion float64 = float64(1000)

/*
Exc016 : What is the sum of the digits of the number 21000
*/
func Exc016() int {
	var number float64 = math.Exp2(expresion)
	var tempStr string = ""
	var sum int = 0

	strNumber := strconv.FormatFloat(number, 'f', 1, 64)
	runes := []rune(strNumber)
	for i := len(strNumber) - 3; i >= 0; i-- {
		tempStr = string(runes[i])
		value, err := strconv.Atoi(tempStr)
		if err != nil {
			return 0
		}
		sum = sum + value
	}

	return sum
}
