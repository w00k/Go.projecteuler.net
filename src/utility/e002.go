package cons

import "fmt"

/*
isEven : evalua si un número es par
param  : int
return : bool
*/
func isEven(x int) bool {
	if x%2 == 0 {
		return true
	}

	return false

}

/*
fibonacci : función recursiva que retorna el valor de fibonacci
param  : int
return : int
*/
func fibonacci(x int) int {
	if x <= 1 {
		return x
	}
	return fibonacci(x-1) + fibonacci(x-2)
}

/*
Exc002 : funcion que retorna la suma de los valores de fibonacci que son pares y menores a 4.000.000
param  : no hay
return : int
*/
func Exc002() int {

	var i, current, sum, max int = 1, 0, 0, 4000000

	for current < max {

		current = fibonacci(i)
		//fmt.Println(i, " ---> ", current)

		if isEven(current) == true && current < max {
			sum = sum + current
			//fmt.Println("sum:", sum)
		}
		i++
	}

	fmt.Println("Exc002 : la suma es ", sum)
	return sum //4613732
}
