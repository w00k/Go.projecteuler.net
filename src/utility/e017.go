package utility

type numberValue struct {
	name  string
	value int
}

var numberStr = map[int]numberValue{
	1:    {"one", 3},
	2:    {"two", 3},
	3:    {"three", 5},
	4:    {"four", 4},
	5:    {"five", 4},
	6:    {"six", 3},
	7:    {"seven", 5},
	8:    {"eight", 5},
	9:    {"nine", 4},
	10:   {"ten", 3},
	11:   {"eleven", 6},
	12:   {"twelve", 6},
	13:   {"thirteen", 8},
	14:   {"fourteen", 8},
	15:   {"fifteen", 7},
	16:   {"sixteen", 7},
	17:   {"seventeen", 9},
	18:   {"eighteen", 8},
	19:   {"nineteen", 8},
	20:   {"twenty", 6},
	30:   {"thirty", 6},
	40:   {"forty", 5},
	50:   {"fifty", 5},
	60:   {"sixty", 5},
	70:   {"seventy", 7},
	80:   {"eighty", 6},
	90:   {"ninety", 6},
	100:  {"hundred", 7},
	1000: {"one thousand", 11},
}

// lessTwentyAndOneNumbers find values between 1 and 21
func lessTwentyAndOneNumbers(index int) int {
	return numberStr[index].value
}

// betweenTwentyAndOneNinetyAndNine find values between 21 and 99
func betweenTwentyAndOneNinetyAndNine(index int) int {
	var decimal int = 0
	var unit int = 0

	if numberStr[index].value != 0 {
		//fmt.Printf("%s \n", numberStr[index].name)
		return numberStr[index].value
	}

	unit = index % 10
	decimal = (index - unit)

	return numberStr[decimal].value + numberStr[unit].value
}

// betweenOneHundredAndOneThousand find values between 99 and 999
func betweenOneHundredAndOneThousand(index int) int {
	var hundred int = 0
	var decimal int = 0
	var and int = 3

	if index%100 == 0 {
		hundred = index / 100
		return numberStr[hundred].value + numberStr[100].value
	}

	decimal = index % 100
	hundred = (index - decimal) / 100

	return numberStr[hundred].value + numberStr[100].value + and + betweenTwentyAndOneNinetyAndNine(decimal)
}

/*
Exc017 : If all the numbers from 1 to 1000 (one thousand) inclusive were written out in words, how many letters would be used?
*/
func Exc017() int {
	var addWords int = 0

	for i := 1; i <= 1000; i++ {
		if i == 1000 || i <= 20 {
			addWords += lessTwentyAndOneNumbers(i)
		}
		if i > 20 && i <= 99 {
			addWords += betweenTwentyAndOneNinetyAndNine(i)
		}
		if i > 99 && i < 1000 {
			addWords += betweenOneHundredAndOneThousand(i)
		}
	}

	return addWords
}
