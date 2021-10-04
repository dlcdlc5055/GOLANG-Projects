package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.InitWindow(900, 980, "2048 Game GOLANG")

	var gd = InitGameData()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		update(&gd)

		rl.ClearBackground(rl.RayWhite)

		if rl.IsKeyReleased(rl.KeyTab) && !(gd.win) {
			reset(&gd)
		}

		if rl.IsKeyReleased(rl.KeySpace) && (gd.win) {
			reset(&gd)
		}

		render(gd)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
