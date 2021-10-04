package main

import rl "github.com/gen2brain/raylib-go/raylib"

func ProcessInput(gd *GameData) {
	if rl.IsKeyPressed(rl.KeyD) {
		if MoveToRight(gd) {
			GenerateNewTile(gd)
		}
	} else if rl.IsKeyPressed(rl.KeyA) {
		if MoveToLeft(gd) {
			GenerateNewTile(gd)
		}
	}
	if rl.IsKeyPressed(rl.KeyW) {
		if MoveToUp(gd) {
			GenerateNewTile(gd)
		}
	} else if rl.IsKeyPressed(rl.KeyS) {
		if MoveToDown(gd) {
			GenerateNewTile(gd)
		}
	}
}
