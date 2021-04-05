package utility

import "fmt"

var arrayNum [][]int = [][]int{
	{75, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{95, 64, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{17, 47, 82, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{18, 35, 87, 10, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{20, 4, 82, 47, 65, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{19, 1, 23, 75, 3, 34, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{88, 2, 77, 73, 7, 63, 67, 0, 0, 0, 0, 0, 0, 0, 0},
	{99, 65, 4, 28, 6, 16, 70, 92, 0, 0, 0, 0, 0, 0, 0},
	{41, 41, 26, 56, 83, 40, 80, 70, 33, 0, 0, 0, 0, 0, 0},
	{41, 48, 72, 33, 47, 32, 37, 16, 94, 29, 0, 0, 0, 0, 0},
	{53, 71, 44, 65, 25, 43, 91, 52, 97, 51, 14, 0, 0, 0, 0},
	{70, 11, 33, 28, 77, 73, 17, 78, 39, 68, 17, 57, 0, 0, 0},
	{91, 71, 52, 38, 17, 14, 91, 43, 58, 50, 27, 29, 48, 0, 0},
	{63, 66, 4, 68, 89, 53, 67, 30, 73, 16, 69, 87, 40, 31, 0},
	{4, 62, 98, 27, 23, 9, 70, 98, 73, 93, 38, 53, 60, 04, 23},
}

var arrayValue []int

func maxValueReturn(a int, b int) int {
	if a < b {
		return b
	}
	return a
}

// left : find the position with the max number
func left(y int, x int) bool {
	var left int = arrayNum[y+1][x]
	var right int = arrayNum[y+1][x+1]
	if y < 13 {
		if maxValueReturn(left+arrayNum[y+2][x], left+arrayNum[y+2][x+1]) < maxValueReturn(right+arrayNum[y+2][x+2], right+arrayNum[y+2][x+1]) {
			return false
		}
		return true
	} else {
		if left < right {
			return false
		}
		return true
	}
}

func loopArray() {
	var x int = 0
	var y int = 0

	arrayValue = append(arrayValue, arrayNum[y][x])

	for y = 0; y < 14; y++ {
		if !left(y, x) {
			x = x + 1
		}
		arrayValue = append(arrayValue, arrayNum[y+1][x])
	}
}

func findMaxSum() int {
	var result int = 0
	for i := 0; i < len(arrayValue); i++ {
		result += arrayValue[i]
	}
	return result
}

/*
Exc018 : Find the maximum total from top to bottom of the triangle below.
75
95 64
17 47 82
18 35 87 10
20 04 82 47 65
19 01 23 75 03 34
88 02 77 73 07 63 67
99 65 04 28 06 16 70 92
41 41 26 56 83 40 80 70 33
41 48 72 33 47 32 37 16 94 29
53 71 44 65 25 43 91 52 97 51 14
70 11 33 28 77 73 17 78 39 68 17 57
91 71 52 38 17 14 91 43 58 50 27 29 48
63 66 04 68 89 53 67 30 73 16 69 87 40 31
04 62 98 27 23 09 70 98 73 93 38 53 60 04 23
*/
func Exc018() int {
	loopArray()

	fmt.Println("len : ", len(arrayValue))
	fmt.Println("arrayValue : ", arrayValue)
	return findMaxSum()
}
