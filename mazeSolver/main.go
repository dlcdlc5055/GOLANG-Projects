package main

import "github.com/gen2brain/raylib-go/raylib"

const bordSize=17

func main() {
	rl.InitWindow(52*17, 52*17, "Mazze Solver")
  var data=generateData()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

    draw(data)
    update(&data)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
