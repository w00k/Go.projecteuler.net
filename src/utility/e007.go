package utility

import (
	"strconv"
)

/*
isPrime : evalua si el número es primo o no
param   : número a evaluar int
return  : bool
*/
func isPrimeExc007(number int) bool {
	middle := number / 2

	for i := 2; i <= middle; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}

/*
Exc007 : obtiene el valor primo 10.001
param  : nil
return : string con el reaultado
*/
func Exc007() string {
	var index, isValue int = 0, 0
	str := ""
	for index = 1; isValue <= 10001; index++ {
		if isPrimeExc007(index) == true {
			isValue++
			str = strconv.Itoa(index)
		}
	}
	return str
}
