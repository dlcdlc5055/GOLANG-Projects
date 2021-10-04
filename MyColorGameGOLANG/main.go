package main

import (
	"math/rand"
	"strconv"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func randInt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}

type Color struct {
	red   uint8
	green uint8
	blue  uint8
}

type GameData struct {
	score          int
	gameDifficulty uint8
	currentCorrect Color
	ColorsToPick   [9]Color
	currentId      int
	state          bool
	offset         int32
	count          int32
}

func generateRandColor() Color {
	var red = randInt(0, 256)
	var green = randInt(0, 256)
	var blue = randInt(0, 256)
	var color = Color{uint8(red), uint8(green), uint8(blue)}
	return color
}

func initGameData() GameData {
	var gd GameData
	gd.currentCorrect = generateRandColor()
	for i := 0; i < 9; i++ {
		gd.ColorsToPick[i] = generateRandColor()
	}
	var buffer = randInt(0, 9)
	gd.currentId = buffer
	gd.ColorsToPick[buffer] = gd.currentCorrect
	gd.gameDifficulty = 2
	gd.score = 0
	gd.state = false
	gd.offset = 40
	gd.count = 0
	return gd
}

func draw(gd GameData) {
	drawHeader(gd)
	drawColorsToPick(gd)
}

func drawHeader(gd GameData) {
	if gd.state == true {
		rl.DrawRectangle(0, 0, 10000, 150, rl.NewColor(gd.currentCorrect.red, gd.currentCorrect.green, gd.currentCorrect.blue, 255))
	} else {
		rl.DrawRectangle(0, 0, 10000, 150, rl.Black)
	}
	rl.DrawText("rgb("+strconv.Itoa(int(gd.currentCorrect.red))+","+strconv.Itoa(int(gd.currentCorrect.green))+","+strconv.Itoa(int(gd.currentCorrect.blue))+")", 20, 25, 100, rl.RayWhite)
	rl.DrawRectangle(0, 150, 1000, 35, rl.Gray)
	if gd.state == true {
		rl.DrawText("Correct!!", 10, 152, 31, rl.RayWhite)
	} else {
		rl.DrawText("score: "+strconv.Itoa(gd.score), 10, 151, 31, rl.RayWhite)
	}
}

func getPickedColor(gd GameData) int {
	for i := 0; i < 3; i++ {
		for j := 0; j < int(gd.gameDifficulty)+1; j++ {
			var mousePos = rl.GetMousePosition()
			var cubeSize int32 = 200
			var offset int32 = gd.offset
			var id int = j*3 + i
			var x_start = 80 + (cubeSize)*int32(i) + offset*int32(i)
			var y_start = 150 + 35 + 40 + (cubeSize)*int32(j) + offset*int32(j)
			var x_end = 80 + (cubeSize)*int32(i) + offset*int32(i) + cubeSize
			var y_end = 150 + 35 + 40 + (cubeSize)*int32(j) + offset*int32(j) + cubeSize
			if x_start < int32(mousePos.X) && y_start < int32(mousePos.Y) && x_end > int32(mousePos.X) && y_end > int32(mousePos.Y) {
				return id + 1
			}
		}
	}
	return -1
}

func drawColorsToPick(gd GameData) {
	for i := 0; i < 3; i++ {
		for j := 0; j < int(gd.gameDifficulty)+1; j++ {
			var clr Color = gd.ColorsToPick[j*3+i]
			if gd.state == false {
				clr = gd.ColorsToPick[j*3+i]
			} else {
				clr = gd.ColorsToPick[gd.currentId]
			}
			var red = clr.red
			var green = clr.green
			var blue = clr.blue
			var color = rl.NewColor(uint8(red), uint8(green), uint8(blue), 255)
			var cubeSize int32 = 200
			var offset int32 = gd.offset
			rl.DrawRectangle(80+(cubeSize)*int32(i)+offset*int32(i), 150+35+40+(cubeSize)*int32(j)+offset*int32(j), cubeSize, cubeSize, color)
		}
	}
}

func checkIfColorEqual(clr1 Color, clr2 Color) bool {
	var red = clr1.red == clr2.red
	var green = clr1.green == clr2.green
	var blue = clr1.blue == clr2.blue
	if red && green && blue {
		return true
	} else {
		return false
	}
}

func generateNewLevelColors(gd *GameData) {
	gd.currentCorrect = generateRandColor()
	for i := 0; i < 9; i++ {
		gd.ColorsToPick[i] = generateRandColor()
	}
	var buffer = randInt(0, 9)
	gd.currentId = buffer
	gd.ColorsToPick[buffer] = gd.currentCorrect
}

func update(gd *GameData) {
	const delay_in_frames = 60 * 1.8
	var picked = getPickedColor(*gd)
	var is_clicked = rl.IsMouseButtonPressed(rl.MouseLeftButton)
	if gd.state == false {
		if is_clicked && picked != -1 {
			var pickedColor = gd.ColorsToPick[picked-1]
			var correctColor = gd.currentCorrect
			var is_Correct = checkIfColorEqual(pickedColor, correctColor)
			if is_Correct {
				gd.state = true
				gd.count = 0
				gd.score += 1
			}
		}
	} else {
		if gd.state == true {
			gd.count = gd.count + 1
			println(gd.count)
			if gd.count > delay_in_frames {
				gd.state = false
				generateNewLevelColors(gd)
			}
		}
	}
}

func main() {
	rl.InitWindow(860, 945, "My Color Game GOlang")

	var gd GameData = initGameData()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		draw(gd)
		update(&gd)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
