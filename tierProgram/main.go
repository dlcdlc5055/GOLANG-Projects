package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.InitWindow(1800, 900, "Tier List Program GOLANG")

	var circles []Circle = GenerateCircles(20)
	var backUpCircles []Circle
	backupCircles(circles, &backUpCircles)
	var selectedCircleIndex = -1
	var isSelectedCircle = false
	var vectOffset = Vector{0, 0}
	var DoneCirclesCount [7]int8
	var DoneCirclesColors [7][7]rl.Color
	print(selectedCircleIndex)
	print(isSelectedCircle)
	print(len(circles))
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		draw(circles, DoneCirclesCount, DoneCirclesColors)

		update(&circles, &selectedCircleIndex, &isSelectedCircle, &vectOffset, &DoneCirclesCount, &DoneCirclesColors, &backUpCircles)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
