package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.InitWindow(880-158, 800-158, "BITLIFE BURGLARY MINI GAME GOLANG WITH ENEMYS")

	var gd GameData = generateGD()

	rl.SetTargetFPS(120)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		draw(gd)
		update(&gd)

		if rl.IsKeyPressed(rl.KeySpace) {
			gd = generateGD()
		}

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
