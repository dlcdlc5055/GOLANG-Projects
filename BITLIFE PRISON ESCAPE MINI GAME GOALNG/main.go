package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.InitWindow(711, 711, "BITLIFE PRISON ESCAPE MINI GAME GOALNG")

	var gd GameData = generateGD()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		draw(&gd)
		update(&gd)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
