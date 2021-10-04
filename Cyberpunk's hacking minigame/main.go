package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.InitWindow(100+85*5+85*3+50, 85*5+100, "Cyberpunk's Hacking Minigame")

	var gd GameData = generateGD()

	rl.SetTargetFPS(60)

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
