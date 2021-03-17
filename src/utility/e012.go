package utility

import (
	"fmt"
	"math"
)

func Exc012() int {
	var maxCount int = 500
	var resultSum int = 0
	var resultDivisors int = 0

	for i := 1; maxCount >= resultDivisors; i++ {
		resultSum = triangleSequenceRule(i)
		resultDivisors = countDivisorsSqrt(resultSum)
	}

	fmt.Println("result sum ", resultSum)
	fmt.Println("count divisors ", resultDivisors)

	return resultSum
}

func triangleSequenceRule(numberTop int) int {
	return numberTop * (numberTop + 1) / 2
}

func countDivisorsSqrt(value int) int {
	var count int = 0
	var valueFloat float64 = float64(value)
	for i := 1; i <= int(math.Sqrt(valueFloat)); i++ {
		if value%i == 0 {
			if value/i == i {
				count += 1
			} else {
				count = count + 2
			}
		}
	}
	return count
}
