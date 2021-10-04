package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(900-2, 902, "Calendar")

	var ap AppData = initAppData()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		update(&ap)

		rl.ClearBackground(rl.Black)

		draw(ap)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
