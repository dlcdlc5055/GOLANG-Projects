package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.InitWindow(1600, 960, "BattleShip Game")

	rl.SetTargetFPS(60)

	var gd = InitGameData()
	gd.bordRed[0][0] = 2
	gd.bordBlue[0][0] = 2

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		update(&gd)

		rl.ClearBackground(rl.RayWhite)

		render(gd)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
