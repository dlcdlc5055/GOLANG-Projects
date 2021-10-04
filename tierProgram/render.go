package main

import rl "github.com/gen2brain/raylib-go/raylib"

func draw(circles []Circle, DoneCirclesCount [7]int8, DoneCirclesColors [7][7]rl.Color) {
	drawBG()
	drawCirclesArray(circles)
}

func drawBG() {
	drawBgBlocks()
	drawBgGrid()
	drawBgLetters()
	rl.DrawRectangle(900+10, 0, 900, 900, rl.LightGray)
}

func drawBgLetters() {
	rl.DrawText("S", 128/2-30+10, 128/2-30, 60, rl.Black)
	rl.DrawText("A", 128/2-30+10, 128/2-30+128, 60, rl.Black)
	rl.DrawText("B", 128/2-30+10, 128/2-30+(128*2), 60, rl.Black)
	rl.DrawText("C", 128/2-30+10, 128/2-30+(128*3), 60, rl.Black)
	rl.DrawText("D", 128/2-30+10, 128/2-30+(128*4), 60, rl.Black)
	rl.DrawText("E", 128/2-30+10, 128/2-30+(128*5), 60, rl.Black)
	rl.DrawText("F", 128/2-30+10, 128/2-30+(128*6), 60, rl.Black)
}

func drawBgGrid() {
	rl.DrawRectangle(0, 0, 10, 900, rl.Black)
	rl.DrawRectangle(128, 0, 10, 900, rl.Black)
	rl.DrawRectangle(900, 0, 10, 900, rl.Black)
	for i := 0; i < 8; i++ {
		rl.DrawRectangle(0, int32(i*128)-4, 900, 10, rl.Black)
	}
}

func drawBgBlocks() {
	rl.DrawRectangle(0, 0, 128, 128, rl.Red)
	rl.DrawRectangle(0, 128, 128, 128, rl.Orange)
	rl.DrawRectangle(0, 128*2, 128, 128, rl.Yellow)
	rl.DrawRectangle(0, 128*3, 128, 128, rl.Green)
	rl.DrawRectangle(0, 128*4, 128, 128, rl.Blue)
	rl.DrawRectangle(0, 128*5, 128, 128, rl.Purple)
	rl.DrawRectangle(0, 128*6, 128, 128, rl.Pink)
}

func drawCirclesArray(circles []Circle) {
	var length = len(circles)
	for i := 0; i < length; i++ {
		drawCircle(circles[i])
	}
}

func drawCircle(circle Circle) {
	if circle.active {
		rl.DrawCircle(circle.x, circle.y, float32(circle.radius+4), rl.Black)
		rl.DrawCircle(circle.x, circle.y, float32(circle.radius), circle.color)
	}
}

func drawWin() {
	rl.DrawRectangle(0, 0, 1800, 900, rl.Black)
	rl.DrawText("Press Space To Reset", 550, 400, 60, rl.RayWhite)
}
