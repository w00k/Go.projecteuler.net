package utility

import (
	"fmt"
	"strconv"
)

func isPalindromo(value int) bool {
	str := strconv.Itoa(value)
	runes := []rune(str)
	tempStr := ""
	//fmt.Println("valor string ", str)
	result := true
	for i := len(str) - 1; i >= 0; i-- {
		tempStr = tempStr + string(runes[i])
		//fmt.Println("temStr --> ", tempStr, " value --> ", str)
	}

	//fmt.Println("tempStr", tempStr,"str ", str)
	if tempStr == str {
		return result
	}
	return !result
}

/*
Exc004  : evalua si es palindromo
param 	: any
return 	: int
*/
func Exc004() int {
	first := 999
	second := 999
	result := 0

	for first >= 100 {
		for second >= 100 {
			if isPalindromo(first*second) && result <= first*second {
				result = first * second
				fmt.Println("---> ", result)
			}
			second--
		}
		second = 999
		first--
	}
	fmt.Println("Exc004 : el valor es ", result)
	return result
}
