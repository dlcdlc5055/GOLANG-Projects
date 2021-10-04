package main

import (
	"math/rand"
	"strconv"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(900, 900, "Rotate Nr Puzzle Game")

	var gd GameData = initGameData()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		draw(gd)
		update(&gd)
		gd.WON = checkForWin(gd)

		//reset
		if rl.IsKeyPressed(rl.KeySpace) {
			gd = initGameData()
		}

		rl.EndDrawing()
	}

	rl.CloseWindow()
}

type GameData struct {
	CURSOR_X  int
	CURSOR_Y  int
	BORD      [4][4]int
	WON       bool
	nrOfMoves int
}

func initGameData() GameData {
	var gd GameData
	gd.CURSOR_X = 1
	gd.CURSOR_Y = 1
	var count = 1
	gd.WON = false
	gd.nrOfMoves = 0
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			gd.BORD[j][i] = count
			count += 1
		}
	}
	randBord(&gd)
	return gd
}

func draw(gd GameData) {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			//draw bord
			var b1 = 0
			rl.DrawCircle(int32(i*225+225/2), int32(j*225+225/2), 100, rl.DarkGreen)
			if gd.BORD[i][j] > 9 && gd.BORD[i][j] != 11 {
				b1 = 10
			} else {
				b1 = -5
			}
			if gd.BORD[i][j] == 1 {
				b1 = -15
			}
			rl.DrawText(strconv.Itoa(gd.BORD[i][j]), int32(i*225+80-b1), int32(j*225+70), 100, rl.White)
		}
	}
	//draw cursor
	rl.DrawCircle(int32(gd.CURSOR_X*225), int32(gd.CURSOR_Y*225), 75, rl.Green)
	//draw win
	if gd.WON == true {
		rl.DrawRectangle(0, 0, 1000, 1000, rl.Black)
		rl.DrawText("You Win. Nr Of Moves "+strconv.Itoa(gd.nrOfMoves), 20, 20, 65, rl.RayWhite)
	}
}

func update(gd *GameData) {
	//process move cursor
	if rl.IsKeyPressed(rl.KeyD) {
		gd.CURSOR_X += 1
	}
	if rl.IsKeyPressed(rl.KeyA) {
		gd.CURSOR_X -= 1
	}
	if rl.IsKeyPressed(rl.KeyW) {
		gd.CURSOR_Y -= 1
	}
	if rl.IsKeyPressed(rl.KeyS) {
		gd.CURSOR_Y += 1
	}
	//corections
	if gd.CURSOR_X < 1 {
		gd.CURSOR_X = 1
	}
	if gd.CURSOR_Y < 1 {
		gd.CURSOR_Y = 1
	}
	if gd.CURSOR_X > 3 {
		gd.CURSOR_X = 3
	}
	if gd.CURSOR_Y > 3 {
		gd.CURSOR_Y = 3
	}
	//rotate pices
	if rl.IsKeyPressed(rl.KeyLeft) {
		rotatePicesInBordLeft(gd)
	}
	if rl.IsKeyPressed(rl.KeyRight) {
		rotatePicesInBordRight(gd)
	}
}

func rotatePicesInBordLeft(gd *GameData) {
	var buffer1 [2][2]int
	var buffer2 [2][2]int
	gd.nrOfMoves += 1
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			buffer1[i][j] = gd.BORD[i-1+gd.CURSOR_X][j-1+gd.CURSOR_Y]
		}
	}
	buffer2[0][0] = buffer1[1][0]
	buffer2[1][0] = buffer1[1][1]
	buffer2[1][1] = buffer1[0][1]
	buffer2[0][1] = buffer1[0][0]
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			gd.BORD[i-1+gd.CURSOR_X][j-1+gd.CURSOR_Y] = buffer2[i][j]
		}
	}
}

func rotatePicesInBordRight(gd *GameData) {
	var buffer1 [2][2]int
	var buffer2 [2][2]int
	gd.nrOfMoves += 1
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			buffer1[i][j] = gd.BORD[i-1+gd.CURSOR_X][j-1+gd.CURSOR_Y]
		}
	}
	buffer2[0][0] = buffer1[0][1]
	buffer2[1][0] = buffer1[0][0]
	buffer2[1][1] = buffer1[1][0]
	buffer2[0][1] = buffer1[1][1]
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			gd.BORD[i-1+gd.CURSOR_X][j-1+gd.CURSOR_Y] = buffer2[i][j]
		}
	}
}

func randBord(gd *GameData) {
	for i := 0; i < randInt(3, 8); i++ {
		c1 := randInt(1, 4)
		c2 := randInt(1, 4)
		type1 := randInt(0, 2)
		gd.CURSOR_X = c1
		gd.CURSOR_Y = c2
		if type1 == 0 {
			rotatePicesInBordLeft(gd)
		} else {
			rotatePicesInBordRight(gd)
		}
	}
	gd.nrOfMoves = 0
	gd.CURSOR_X = 2
	gd.CURSOR_Y = 2
}

func randInt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}

func checkForWin(gd GameData) bool {
	var count = 1
	gd.WON = false
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if gd.BORD[j][i] != count {
				return false
			}
			count += 1
		}
	}
	return true
}
