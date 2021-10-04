package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.InitWindow(1200, 900, "Columns Game GOLANG")

	var gd GameData = generateGameData()

	var paused = true

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		draw(gd, paused)

		if !paused {
			update(&gd)
		}

		if rl.IsKeyPressed(rl.KeyTab) {
			paused = !paused
		}

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
