package utility

import "fmt"

const gridSize int = 20

func Exc015() int {
	var paths int = 1
	fmt.Printf("The grid has %d to %d", gridSize, gridSize)

	for i := 0; i < gridSize; i++ {
		paths *= (2 * gridSize) - i
		paths /= i + 1
	}

	return paths
}
