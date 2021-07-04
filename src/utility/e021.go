package utility

func divisorSum(n int) int {
	sum := 0
	for i := 1; i < n; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	return sum
}

func isAmicable(n int) bool {
	m := divisorSum(n)
	return m != n && divisorSum(m) == n
}

// index > 1;
func Exc021() int {
	sum := 0
	for i := 1; i < 10000; i++ {
		if isAmicable(i) {
			sum += i
		}
	}
	return sum
}
