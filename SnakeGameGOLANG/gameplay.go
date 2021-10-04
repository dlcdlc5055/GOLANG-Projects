package main

import "github.com/gen2brain/raylib-go/raylib"

func ProcessInput(gd *GameData) {
	if rl.IsKeyPressed(rl.KeyD) && gd.direction != left {
		gd.direction = right
		return
	}
	if rl.IsKeyPressed(rl.KeyA) && gd.direction != right {
		gd.direction = left
		return
	}
	if rl.IsKeyPressed(rl.KeyW) && gd.direction != down {
		gd.direction = up
		return
	}
	if rl.IsKeyPressed(rl.KeyS) && gd.direction != up {
		gd.direction = down
		return
	}
}
