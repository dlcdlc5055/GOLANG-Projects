package main

import "github.com/gen2brain/raylib-go/raylib"

func main() {

	rl.InitWindow(1600, 900, "ArreySort Visualizer")

	var array [400]int
	var b1, b2, active int
	var paused bool = false
	array = ArrayGenerate()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.DrawFPS(10, 10)

		if !paused {
			MultiStageUpdate(&b1, &b2, &active, &array, 36)
		}

		if rl.IsKeyPressed(rl.KeyTab) {
			paused = !paused
		}

		rl.ClearBackground(rl.DarkGray)

		renderArray(b1, b2, array)

		if CheckIfSorted(&array) {
			reset(&b1, &b2, &active, &array)
		}

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
