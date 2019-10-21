package cons

import "fmt"

/*
Exc001 : Función que retorna la suma de los multiplos de 3 y 5, que son menores de 1000.
Parametros :  de entrada no son necesarios
Retorno : int
*/
func Exc001() int {

	var three, five, result int = 3, 5, 0

	for i := 1000 - 1; i >= 0; i-- {

		if i%three == 0 {
			result = result + i
		} else if i%five == 0 {
			result = result + i
		}

	}

	fmt.Println("Exc001 : la suma corresponde a ", result)
	return result
}
