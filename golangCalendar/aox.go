package main

import (
	"math"
	"time"
)

func calcDiference(val1, val2 int) int {
	if val2 > val1 {
		return val2 - val1
	} else if val2 < val1 {
		return val1 - val2
	} else {
		return 0
	}
}

func calcDistance(x1, y1, x2, y2 int) int {
	var dX = calcDiference(x1, x2)
	var dY = calcDiference(y1, y2)
	return int(math.Sqrt(float64(dX*dX + dY*dY)))
}

func checkIfToday(date time.Time) bool {
	if int(date.Month()) == int(time.Now().Month()) && date.Year() == time.Now().Year() && time.Now().Day() == date.Day() {
		return true
	}
	return false
}

func getFirstDayOfMonthId(ap AppData) int {
	var date = time.Date(ap.year, time.Month(ap.monthId+1), 1, 0, 0, 0, 0, time.UTC)
	return int(date.Weekday())
}
