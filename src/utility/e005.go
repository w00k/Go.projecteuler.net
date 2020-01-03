package utility

import "fmt"

type exp struct {
	base int
	expo int
}

/*
isPrime : verifica si un número es primo o no
param   : int
return  : bool
*/
func isPrimeExc005(number int) bool {
	middle := number / 2
	for i := 2; i <= middle; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}

/*
getBase : obtiene el número base para cada número
param   : se ingresa el número a consultar
return  :
*/
func getBase(number int) int {
	var auxNumber, base int = 1, 1

	for base = 2; base < number; base++ {
		auxNumber = 1
		for exp := 1; auxNumber < number; exp++ {
			auxNumber = base * auxNumber
			//fmt.Println("base: ", base, " - exp: ", "auxNumber: ", auxNumber, "number: ", number)
			if auxNumber == number {
				//fmt.Println("iguales ---> ", "auxNumber: ", auxNumber, "number: ", number)
				return base
			}
		}
	}

	return 1
}

/*
Exc005 :
param  : no hay
return : maximo como un divisor entre los valores del 1 al 20
*/
func Exc005() int {
	number := 1
	for i := 1; i <= 20; i++ {
		if !isPrimeExc005(i) {
			fmt.Println("not prime ", i)
			number = number * getBase(i)
		} else {
			fmt.Println("is prime", i)
			number = number * i
		}
	}

	fmt.Println("Exc005 ===> ", number)
	return number
}
