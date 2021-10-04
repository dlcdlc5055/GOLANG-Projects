package main

import "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.InitWindow(900, 950, "Snake")

	var gd = GameData{}
	initGD(&gd)

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		update(&gd)

		rl.ClearBackground(rl.Black)

		render(gd)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
