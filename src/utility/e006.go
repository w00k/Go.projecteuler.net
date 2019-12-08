package cons

import (
	"fmt"
	"strconv"
)

/*
Exc006NaturalNumbers : return the sum for 1 to 100 and square
param				 : any
return 				 : int64
*/
func Exc006NaturalNumbers() int64 {
	value := 100
	var sum int64 = 0
	for i := 0; i <= value; i++ {
		sum = sum + int64(i)
	}
	return sum * sum
}

/*
Exc006NaturalNumbersSquare : return the sum of the square from 1 to 100
param				 : any
return 				 : int64
*/
func Exc006NaturalNumbersSquare() int64 {
	value := 100
	var sum int64 = 0
	for i := 0; i <= value; i++ {
		sum = sum + int64(i*i)
	}
	return sum
}

/*
Exc006 : Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.
param  : any
return : string
*/
func Exc006() string {
	naturalNumbers := Exc006NaturalNumbers()
	naturalNumbersSquare := Exc006NaturalNumbersSquare()

	fmt.Println("Exc006NaturalNumbers : " + strconv.FormatInt(naturalNumbers, 10))
	fmt.Println("Exc006NaturalNumbersSquare : " + strconv.FormatInt(naturalNumbersSquare, 10))
	return strconv.FormatInt(naturalNumbers-naturalNumbersSquare, 10)
}
