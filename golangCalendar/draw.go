package main

import (
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const navbarHeight = 75
const calendarHeight = 900 - 75 - 55

func draw(ap AppData) {
	drawDays(ap)
	drawTopUI(ap)
}

func drawTopUI(ap AppData) {
	var mousePos = rl.GetMousePosition()
	rl.DrawRectangle(0, 0, 10000, 75, rl.DarkGreen)
	var dateText = ap.monthName + " " + strconv.Itoa(ap.year)
	var posVect = rl.Vector2{10, 10}
	rl.DrawTextEx(ap.font, dateText, posVect, 55, 5, rl.RayWhite)
	rl.DrawCircle(900-50, 38, 30, rl.Black)
	rl.DrawCircle(900-50-100, 38, 30, rl.Black)
	rl.DrawCircle(900-50, 38, 25, rl.Green)
	rl.DrawCircle(900-50-200, 38, 30, rl.Black)
	rl.DrawCircle(900-50-200, 38, 25, rl.Green)
	if calcDistance(900-50, 38, int(mousePos.X), int(mousePos.Y)) < 25 {
		rl.DrawCircle(900-50, 38, 25, rl.Lime)
	}
	if calcDistance(900-50, 38, int(mousePos.X), int(mousePos.Y)) < 25 && rl.IsMouseButtonDown(rl.MouseLeftButton) {
		rl.DrawCircle(900-50, 38, 25, rl.RayWhite)
	}
	rl.DrawCircle(900-50-100, 38, 25, rl.Green)
	if calcDistance(900-50-100, 38, int(mousePos.X), int(mousePos.Y)) < 25 {
		rl.DrawCircle(900-50-100, 38, 25, rl.Lime)
	}
	if calcDistance(900-50-100, 38, int(mousePos.X), int(mousePos.Y)) < 25 && rl.IsMouseButtonDown(rl.MouseLeftButton) {
		rl.DrawCircle(900-50-100, 38, 25, rl.RayWhite)
	}
	if calcDistance(900-50-200, 38, int(mousePos.X), int(mousePos.Y)) < 25 {
		rl.DrawCircle(900-50-200, 38, 25, rl.Lime)
	}
	if calcDistance(900-50-200, 38, int(mousePos.X), int(mousePos.Y)) < 25 && rl.IsMouseButtonDown(rl.MouseLeftButton) {
		rl.DrawCircle(900-50-200, 38, 25, rl.RayWhite)
	}

	rl.DrawRectangle(0, 75, 1000, 55, rl.Gray)
	for i := 0; i < 7; i++ {
		var x int32 = int32(i * (900 / 7))
		var y int32 = 75
		var dayStr string = ap.DaysOfWeekString[i][0:3]
		if i != 0 {
			rl.DrawRectangle(x, y, 2, 55, rl.Black)
		}
		rl.DrawText(dayStr, x+16, y+7, 44, rl.Black)
	}
	rl.DrawRectangle(0, 75+55, 1000, 2, rl.Black)
}

func drawDays(ap AppData) {
	if len(ap.days) == 0 {
		return
	}
	const day_size_width = 900 / 7
	var day_size_height int32 = int32(calendarHeight / (len(ap.days) / 7))
	var weeksInCalendar int32 = int32(len(ap.days) / 7)
	//draw days
	for i := 0; i < len(ap.days); i++ {
		var x = int32(i % 7)
		var y = int32(i / 7)
		var colorBG = rl.Gray
		var colorText = rl.Black
		if ap.days[i].active == false {
			colorText = rl.DarkGray
		}
		if ap.days[i].active && !ap.days[i].today {
			colorBG = rl.Beige
		} else if ap.days[i].active && ap.days[i].today {
			colorBG = rl.DarkGreen
			colorText = rl.RayWhite
		}
		rl.DrawRectangle(x*day_size_width, y*day_size_height+75+55, day_size_width, int32(day_size_height), colorBG)
		rl.DrawText(strconv.Itoa(ap.days[i].number), x*day_size_width+10, y*day_size_height+75+55+10, 43, colorText)
	}
	//draw grid
	for i := 0; i < 8; i++ {
		rl.DrawRectangle(int32(day_size_width*i), int32(0), 2, 10000, rl.Black)
	}
	for i := 0; i < int(weeksInCalendar)+1; i++ {
		rl.DrawRectangle(int32(0), int32(i)*day_size_height+55+75, 10000, 2, rl.Black)
	}
}
