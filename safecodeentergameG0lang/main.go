package main

import (
	"math"
	"math/rand"
	"strconv"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func checkIfWin(input [3]int, combination [3]int) bool {
	if input[0] == combination[0] && input[1] == combination[1] && input[2] == combination[2] {
		return true
	}
	return false
}

func generateCombination() [3]int {
	var combination [3]int
	rand.Seed(time.Now().UTC().UnixNano())
	combination[0] = rand.Intn(10)
	combination[1] = rand.Intn(10)
	combination[2] = rand.Intn(10)
	for true {
		if combination[0] == combination[1] || combination[1] == combination[2] || combination[0] == combination[2] {
			combination[0] = rand.Intn(10)
			combination[1] = rand.Intn(10)
			combination[2] = rand.Intn(10)
		} else {
			break
		}
	}
	return combination
}

func drawCombination(combination [3]int) {
	rl.DrawRectangle(15, 15, 60*6-20, 70, rl.DarkGray)
	rl.DrawRectangle(20, 20, 60*6-30, 60, rl.Black)
	var t1 = strconv.Itoa(combination[0] * 10)
	var t2 = strconv.Itoa(combination[1] * 10)
	var t3 = strconv.Itoa(combination[2] * 10)
	if t1 == "0" {
		t1 += "0"
	}
	if t2 == "0" {
		t2 += "0"
	}
	if t3 == "0" {
		t3 += "0"
	}
	rl.DrawText(t1+" | "+t2+" | "+t3, 30, 24, 58, rl.RayWhite)
}

func drawMarks(marks [3]float32, marked [3]bool, color rl.Color) {
	for i := 0; i < 3; i++ {
		var buffer = float64(marks[i] - 90)
		var radians = buffer * (math.Pi / 180)
		buffer = radians
		if marked[i] {
			rl.DrawCircle(int32(400+270*math.Cos(float64(buffer))), int32(400+270*math.Sin(float64(buffer))), 10, color)
		}
	}
}

func drawWin() {
	rl.DrawRectangle(0, 0, 1000, 1000, rl.Black)
	rl.DrawText("YOU WIN!!", 50, 50, 50, rl.RayWhite)
	rl.DrawText("PRESS SPACE TO RESET!!", 50, 200, 50, rl.RayWhite)
}

func drawLose() {
	rl.DrawRectangle(0, 0, 1000, 1000, rl.Black)
	rl.DrawText("YOU LOSE!!", 50, 50, 50, rl.RayWhite)
	rl.DrawText("PRESS SPACE TO RESET!!", 50, 200, 50, rl.RayWhite)
}

func reset(safeRotation *float32, direction *bool, combination *[3]int, marks *[3]float32, marked *[3]bool, input *[3]int, count *int, r1 bool, canMove *bool) {
	*safeRotation = 0
	*direction = false
	if r1 {
		*combination = generateCombination()
	}
	for i := 0; i < 3; i++ {
		marked[i] = false
		input[i] = -1
	}
	*count = 0
	*canMove = true
}

func main() {
	rl.InitWindow(800-25, 800-50, "Safe lock golang")

	go rl.SetTargetFPS(240)

	var canMove = true
	var safeTexture = rl.LoadTexture("img/safe_lock.png")
	var safeVect rl.Vector2
	safeVect.X = 25
	safeVect.Y = 0
	var safeRotation float32 = 0
	const rotationSpeed float32 = 90
	var direction = false
	var combination [3]int = generateCombination()
	var marks [3]float32
	var marked [3]bool
	var input [3]int
	var count = 0

	go reset(&safeRotation, &direction, &combination, &marks, &marked, &input, &count, false, &canMove)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.DrawTexturePro(safeTexture, rl.NewRectangle(0, 0, 750, 800), rl.NewRectangle(400, 400, 750, 800), rl.NewVector2(750/2, 400), safeRotation, rl.RayWhite)
		rl.DrawRectangle(400-4, 400-320, 8, 30, rl.Yellow)

		drawMarks(marks, marked, rl.DarkGreen)

		drawCombination(combination)

		if rl.IsKeyPressed(rl.KeyS) && canMove && !rl.IsKeyDown(rl.KeyW) {
			go reset(&safeRotation, &direction, &combination, &marks, &marked, &input, &count, false, &canMove)
		}

		if rl.IsKeyDown(rl.KeyW) && canMove {
			if direction == false {
				safeRotation += rotationSpeed * rl.GetFrameTime()
			} else {
				safeRotation -= rotationSpeed * rl.GetFrameTime()
			}
			if safeRotation > 360 {
				safeRotation = 0
			}
			if safeRotation < 0 {
				safeRotation = 360
			}
			for i := 0; i < 3; i++ {
				if direction == false {
					marks[i] += rotationSpeed * rl.GetFrameTime()
				} else {
					marks[i] -= rotationSpeed * rl.GetFrameTime()
				}
			}
		} else {
			for i := 0; i < 3; i++ {
				if marked[i] {
					marks[i] = float32(safeRotation + float32(input[i])*36)
				}
			}
			if int(safeRotation)%36 != 0 {
				var buffer = int(safeRotation) % 36
				if buffer < 17 {
					safeRotation = float32(int(int(safeRotation) - (buffer)))
				} else {
					safeRotation = float32(int(int(safeRotation) + (buffer - 16)))
				}
			}
			var in = 10 - int(safeRotation/36)
			if in == 10 {
				in = 0
			}
			if count < 3 && rl.IsKeyPressed(rl.KeySpace) {
				var can = true
				for i := 0; i < 3; i++ {
					if in == input[i] {
						can = false
					}
				}
				if can {
					input[count] = in
					marked[count] = true
					marks[count] = 0
					count += 1
					direction = !direction
				}
			} else if count >= 3 {
				if checkIfWin(input, combination) {
					canMove = false
					drawWin()
				} else {
					drawLose()
					canMove = false
				}
				if rl.IsKeyPressed(rl.KeySpace) {
					go reset(&safeRotation, &direction, &combination, &marks, &marked, &input, &count, true, &canMove)
				}
			}
		}

		rl.EndDrawing()
	}

	go rl.CloseWindow()
}
