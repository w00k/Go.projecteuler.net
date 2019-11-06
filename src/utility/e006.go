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
addValues : retorna la suma de 1 a 100
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
addValuesSquares : obtiene la suma de los cuadrados de 1 al 100
param            : nil
return           : value int
*/
func addValuesSquares() int {
	value := 0
	for index := 1; index < 101; index++ {
		value = value + squares(index)
	}
	return value
}

/*
Exc006 : obtiene la diferencia entre la suma al cuadrado del 1 al 100 a la suma de los cuadrados de 1 al 100
param  : nil
return : retorna la diferencia entre la suma al cuadrado del 1 al 100 a la suma de los cuadrados de 1 al 100 
*/
func Exc006() int {
	addValue := squares(addValues())
	addSquaresValue := addValuesSquares()
	fmt.Println("values 1 + 2 + 3 + ... + 101", addValue )
	fmt.Println("return addValue : ", addValue)
	fmt.Println("return addSquaresValue : ", addSquaresValue)
	return addValue - addSquaresValue
}
