package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(800, 800, "Fractal Generator")

	var generated = false
	var image [800][800]uint8

	rl.SetTargetFPS(60)

	go generateFractal(&image, &generated)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		if generated == true {
			for i := 0; i < 800; i++ {
				for j := 0; j < 800; j++ {
					var color = rl.Color{image[i][j], 0, 0, 255}
					rl.DrawRectangle(int32(i), int32(j), 1, 1, color)
				}
			}
		}

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
