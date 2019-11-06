package cons

import "fmt"

/*
squares : retirno el valor de un número al cuadrado
param   : número base
return  : number int
*/
func squares(number int) int {
	return number * number
}

/*
addValues : retorna la suma de 1 a 101
param     : nil
return    : value int
*/
func addValues() int {
	value := 0
	for index := 1; index < 101; index++ {
		value = value + index
	}
	return value
}

/*

*/
func addValuesSquares() int {
	value := 0
	for index := 1; index < 101; index++ {
		value = value + squares(index)
	}
	return value
}
/*
Exc006 :
*/
func Exc006() int {
	addValue := squares(addValues())
	addSquaresValue := addValuesSquares()
	fmt.Println("values 1 + 2 + 3 + ... + 101", addValue )
	fmt.Println("return addValue : ", addValue)
	fmt.Println("return addSquaresValue : ", addSquaresValue)
	return addValue - addSquaresValue
}
