package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func draw(gd GameData) {
	drawPlayerPos(gd)
	drawCorrectPos(gd)
	drawBtns()
}

func drawMagnet(c1, c2 rl.Color, x, y int32) {
	rl.DrawRectangle(x, y, 80, 80, c1)
	rl.DrawRectangle(x+80, y, 80, 80, c2)
}

func drawPlayerPos(gd GameData) {
	var c1, c2 rl.Color
	for i := 0; i < 8; i++ {
		if gd.player[i] == false {
			c1 = rl.Blue
			c2 = rl.Red
		} else {
			c1 = rl.Red
			c2 = rl.Blue
		}
		drawMagnet(c1, c2, 300, 50+int32(i)*90)
	}
}

func drawCorrectPos(gd GameData) {
	var c1, c2 rl.Color
	for i := 0; i < 8; i++ {
		if gd.correct[i] == false {
			c1 = rl.Blue
			c2 = rl.Red
		} else {
			c1 = rl.Red
			c2 = rl.Blue
		}
		drawMagnet(c1, c2, 700, 50+int32(i)*90)
	}
}

func drawBtns() {
	var buffer bool
	for i := 0; i < 7; i++ {
		var color = rl.Gray
		if calculateDistance(200, 50+int32(i)*90+80, int32(rl.GetMousePosition().X), int32(rl.GetMousePosition().Y)) < 30 {
			buffer = true
		} else {
			buffer = false
		}
		if buffer {
			color = rl.DarkGreen
		} else {
			color = rl.Gray
		}
		rl.DrawCircle(200, 50+int32(i)*90+80, 30, rl.DarkGreen)
		rl.DrawCircle(200, 50+int32(i)*90+80, 20, color)
		rl.DrawCircle(200, 50+int32(i)*90+80, 15, rl.DarkGreen)
	}
}
