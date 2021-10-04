package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.InitWindow(1200, 800, "Simple Magnets Game GOLANG")

	var gd = generateGameData()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.Gray)

		draw(gd)

		update(&gd)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
