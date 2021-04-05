package utility

import (
	"fmt"
	"math/big"
	"strconv"
)

func loopTop() {
	result := big.NewInt(1)
	one := big.NewInt(1)
	hundred := big.NewInt(100)
	for i := big.NewInt(2); i.Cmp(hundred) <= 0; i.Add(i, one) {
		result.Mul(result, i)
	}
	fmt.Printf("result ::: %d \n", result)
}

// Find the sum of the digits in the number 100!
func Exc020() int {

	result := big.NewInt(1)
	one := big.NewInt(1)
	hundred := big.NewInt(100)
	strNumber := ""
	resultInt := 0

	for i := big.NewInt(2); i.Cmp(hundred) <= 0; i.Add(i, one) {
		result.Mul(result, i)
	}

	fmt.Printf("result ::: %d \n", result)

	strNumber = result.String()
	for _, x := range strNumber {
		fmt.Println(string(x))
		xint, _ := strconv.Atoi(string(x))
		resultInt += xint
	}

	fmt.Println("End")
	return resultInt
}
