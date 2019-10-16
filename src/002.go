package main

import "fmt"

func isEven(x int) bool {
	if x%2 == 0 {
		return true
	} else {
		return false
	}
}

func fibonacci(x int) int {
	if x <= 1 {
		return x
	}
	return fibonacci(x-1) + fibonacci(x-2)
}

func main() {

	var i, current, sum, max int = 1, 0, 0, 4000000

	for current < max {

		current = fibonacci(i)
		fmt.Println(i, " ---> ", current)

		if isEven(current) == true && current < max {
			sum = sum + current
			fmt.Println("sum:", sum)
		}
		i++
	}

	fmt.Println("valor es:", sum)
	//4613732
}
