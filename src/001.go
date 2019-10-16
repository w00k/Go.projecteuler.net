package main

import "fmt"

func main() {

	var three, five, result int = 3, 5, 0

	for i := 1000 - 1; i >= 0; i-- {

		if i%three == 0 {
			result = result + i
		} else if i%five == 0 {
			result = result + i
		}

	}

	fmt.Println("el valor es ", result)

}
