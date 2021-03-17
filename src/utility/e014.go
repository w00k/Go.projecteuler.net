package utility

import "fmt"

func Exc014() int {
	var maxIteration int = 0
	var indexValue int = 0
	for i := 999999; 1 <= i; i-- {
		countIteration := countTerms(i)
		if maxIteration < countIteration {
			maxIteration = countIteration
			indexValue = i
		}
	}
	fmt.Println("count iteration : ", maxIteration)
	fmt.Println("index max : ", indexValue)
	return indexValue
}

func countTerms(number int) int {
	var iterativeNumber int = number
	var count int = 0
	for iterativeNumber > 1 {
		if isEven(iterativeNumber) {
			iterativeNumber = iterativeNumber / 2
		} else {
			iterativeNumber = (3 * iterativeNumber) + 1
		}
		count += 1
	}
	return count
}

/* Exc002
func isEven(number int) bool {
	if number%2 == 0 {
		return true
	}
	return false
}
*/
