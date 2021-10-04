package main

import rl "github.com/gen2brain/raylib-go/raylib"

func update(ap *AppData) {
	var mousePos = rl.GetMousePosition()
	if (calcDistance(900-50, 38, int(mousePos.X), int(mousePos.Y)) < 25 && rl.IsMouseButtonPressed(rl.MouseLeftButton)) || (rl.IsKeyPressed(rl.KeyD) || rl.IsKeyPressed(rl.KeyRight)) {
		setDateValuesForOffsetType(ap, true)
	}
	if (calcDistance(900-50-200, 38, int(mousePos.X), int(mousePos.Y)) < 25 && rl.IsMouseButtonPressed(rl.MouseLeftButton)) || (rl.IsKeyPressed(rl.KeyA) || rl.IsKeyPressed(rl.KeyLeft)) {
		setDateValuesForOffsetType(ap, false)
	}
	if (calcDistance(900-50-100, 38, int(mousePos.X), int(mousePos.Y)) < 25 && rl.IsMouseButtonPressed(rl.MouseLeftButton)) || (rl.IsKeyPressed(rl.KeyF5)) {
		*ap = initAppData()
	}
}
