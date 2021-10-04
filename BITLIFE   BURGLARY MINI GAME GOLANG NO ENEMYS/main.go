package main

import rl "github.com/gen2brain/raylib-go/raylib"

const blockSize = 75

func main() {
	rl.InitWindow(825+75, 825+50, "BITLIFE BURGLARY MINI GAME GOLANG")

	var gd GameData = generateGD()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		if rl.IsKeyReleased(rl.KeySpace) {
			gd = generateGD()
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		if !gd.poused {
			update(&gd)
		}
		draw(gd)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
