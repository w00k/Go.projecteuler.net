package utility

import "fmt"

//There exists exactly one Pythagorean triplet for which a + b + c = 1000.
//Find the product abc.

//validate sum
func validatePythagorean(a int, b int, c int) bool {
	sum := (a * a) + (b * b)
	add := c * c
	if sum == add {
		allsum := a + b + c
		if allsum == 1000 {
			fmt.Println("a :", a)
			fmt.Println("b :", b)
			fmt.Println("c :", c)
			fmt.Println("result : ", (a * b * c))
			return true
		}
		return false
	}
	return false
}

//Exc009 function to run the solution
func Exc009() bool {
	var a, b, c int
	max := 500
	for a = 1; a < max; a++ {
		for b = 1; b < max; b++ {
			for c = 1; c < max; c++ {
				if validatePythagorean(a, b, c) {
					return true
				}
			}
		}
	}
	return false
}
