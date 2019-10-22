package cons

import "fmt"

const constant = 600851475143

/*
isPrime :
*/
func isPrime(number int) bool {
	middle := number / 2

	for i := 2; i <= middle; i++ {
		if number%i == 0 {
			return false
		}
	}

	return true
}

/*
Exc003 : función que retorna los números primos del número constante
param  : no hay
return : int
*/
func Exc003() int {
	middle := constant / 2
	temporalContant := constant
	auxiliar := 0
	auxiliarBool := false
	fmt.Println("600851475143 ==> ", constant, " ---> ", middle)

	for i := 2; i <= middle && auxiliarBool == false; i++ {
		if temporalContant%i == 0 {
			auxiliarDiv := temporalContant / i
			if isPrime(auxiliarDiv) && auxiliar < auxiliarDiv {
				auxiliar = temporalContant / i
				fmt.Println("valor auxiliar ", auxiliar)
				auxiliarBool = !auxiliarBool
			}
		}
	}
	return auxiliar
}
