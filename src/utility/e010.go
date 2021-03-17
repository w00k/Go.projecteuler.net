package utility

import (
	"fmt"
	"math"
)

/*
The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

Find the sum of all the primes below two million.
*/

/*
IsPrimeSqrt : evalua si el número es primo o no
param   : número a evaluar int
return  : bool
*/
func isPrimeSqrt(value int) bool {
	for i := 2; i <= int(math.Floor(math.Sqrt(float64(value)))); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}

//Exc010 function to run the solution
func Exc010() int {
	sum := 0
	for i := 2; i <= 2000000; i++ {
		if isPrimeSqrt(i) {
			sum = sum + i
		}
	}
	fmt.Println("suma total : ", sum)
	return sum
}
