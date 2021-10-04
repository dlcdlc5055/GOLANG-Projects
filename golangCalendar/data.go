package main

import (
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type AppData struct {
	offset           int
	days             []Day
	date             time.Time
	monthsString     []string
	DaysOfWeekString []string
	monthName        string
	monthId          int
	year             int
	dayOfMonth       int
	monthLenght      int
	WeekDayName      string
	dayId            int
	lastMonthLenght  int
	font             rl.Font
}

type Day struct {
	active bool
	number int
	today  bool
}

func initAppData() AppData {
	var ap AppData
	ap.offset = 0
	ap.date = time.Now()
	//init days of month
	ap.monthsString = append(ap.monthsString, "January")
	ap.monthsString = append(ap.monthsString, "February")
	ap.monthsString = append(ap.monthsString, "March")
	ap.monthsString = append(ap.monthsString, "April")
	ap.monthsString = append(ap.monthsString, "May")
	ap.monthsString = append(ap.monthsString, "June")
	ap.monthsString = append(ap.monthsString, "July")
	ap.monthsString = append(ap.monthsString, "August")
	ap.monthsString = append(ap.monthsString, "September")
	ap.monthsString = append(ap.monthsString, "October")
	ap.monthsString = append(ap.monthsString, "November")
	ap.monthsString = append(ap.monthsString, "December")
	//init days of week
	ap.DaysOfWeekString = append(ap.DaysOfWeekString, "Sunday")
	ap.DaysOfWeekString = append(ap.DaysOfWeekString, "Monday")
	ap.DaysOfWeekString = append(ap.DaysOfWeekString, "Tuesday")
	ap.DaysOfWeekString = append(ap.DaysOfWeekString, "Wednesday")
	ap.DaysOfWeekString = append(ap.DaysOfWeekString, "Thursday")
	ap.DaysOfWeekString = append(ap.DaysOfWeekString, "Friday")
	ap.DaysOfWeekString = append(ap.DaysOfWeekString, "Saturday")
	//set date values
	ap.monthName = ap.monthsString[ap.date.Month()-1]
	ap.monthId = int(ap.date.Month()) - 1
	ap.year = ap.date.Year()
	ap.dayOfMonth = ap.date.Day()
	ap.monthLenght = EndOfMonth(ap.date).Day()
	ap.dayId = int(ap.date.Weekday())
	ap.WeekDayName = ap.DaysOfWeekString[ap.dayId]
	var buffer = time.Date(ap.year, time.Month(ap.monthId), ap.dayOfMonth, 0, 0, 0, 0, time.UTC)
	ap.lastMonthLenght = EndOfMonth(buffer).Day()
	//reset date value
	ap.font = rl.LoadFont("resources/fonts/mecha.png")
	ap.date = time.Now()
	//generate days
	ap.days = genearateDays(ap)
	return ap
}

func setDateValuesForOffsetType(ap *AppData, value bool) {
	//todo this
	if value == true {
		ap.offset += 1
	} else {
		ap.offset -= 1
	}
	ap.date = time.Now()
	ap.date = ap.date.AddDate(0, ap.offset, 0)
	ap.monthName = ap.monthsString[ap.date.Month()-1]
	ap.monthId = int(ap.date.Month()) - 1
	ap.year = ap.date.Year()
	ap.dayOfMonth = ap.date.Day()
	ap.monthLenght = EndOfMonth(ap.date).Day()
	ap.dayId = int(ap.date.Weekday())
	ap.WeekDayName = ap.DaysOfWeekString[ap.dayId]
	var buffer = time.Date(ap.year, time.Month(ap.monthId), ap.dayOfMonth, 0, 0, 0, 0, time.UTC)
	ap.lastMonthLenght = EndOfMonth(buffer).Day()
	ap.days = genearateDays(*ap)
}

func EndOfMonth(date time.Time) time.Time {
	return date.AddDate(0, 1, -date.Day())
}

func genearateDays(ap AppData) []Day {
	var days []Day
	var days_lenght = 0
	//prev days
	for i := 0; i < getFirstDayOfMonthId(ap); i++ {
		var day Day = Day{false, ap.lastMonthLenght - getFirstDayOfMonthId(ap) + i + 1, false}
		days = append(days, day)
		days_lenght += 1
	}
	//month days
	for i := 1; i <= int(ap.monthLenght); i++ {
		var date = time.Date(ap.year, time.Month(ap.monthId+1), i, 0, 0, 0, 0, time.UTC)
		var is_today = checkIfToday(date)
		var day Day = Day{true, i, is_today}
		days = append(days, day)
		days_lenght += 1
	}
	//next days
	var count = 0
	for true {
		count++
		if days_lenght%7 == 0 {
			break
		}
		days_lenght += 1
		var day Day = Day{false, count, false}
		days = append(days, day)
	}
	//next days
	return days
}
