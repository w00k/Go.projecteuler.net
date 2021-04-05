package utility

import (
	"time"
)

// add 1 day to t variable, start on 1902-01-01
func addDay(t time.Time) time.Time {
	return t.Add(time.Hour * 24)
}

// How many Sundays fell on the first of the month during the twentieth century (1 Jan 1901 to 31 Dec 2000)?
func Exc019() int {
	str := "1901-01-01T10:45:00.371Z"
	endStr := "2000-12-31T10:45:00.371Z"
	t, _ := time.Parse(time.RFC3339, str)
	end, _ := time.Parse(time.RFC3339, endStr)
	count := 0

	for t.Unix() <= end.Unix() {
		t = addDay(t)
		if t.Day() == 1 && t.Weekday().String() == "Sunday" {
			count += 1
		}
	}

	return count
}
