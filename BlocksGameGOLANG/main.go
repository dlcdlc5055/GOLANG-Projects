package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.InitWindow(1600, 900, "BlockGame")

	var gd = GameData{}
	InitGameData(&gd)

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		render(gd)

		if rl.IsKeyReleased(rl.KeyTab) {
			reset(&gd)
		}

		update(&gd)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
