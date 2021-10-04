package main

import rl "github.com/gen2brain/raylib-go/raylib"

func update(gd *GameData) {
	var win = checkForWin(*gd)
	var lose = checkForLose(*gd)

	if checkIfInSelected(*gd) && rl.IsMouseButtonPressed(rl.MouseLeftButton) && gd.bordSubData[getMouseX()][getMouseY()] == true && !win && !lose {
		var buffer = gd.bord[getMouseX()][getMouseY()]
		if buffer == gd.correct[gd.bufferIndex] {
			gd.buffer[gd.bufferIndex] = buffer
			gd.bufferIndex += 1
			gd.corrects += 1
		} else {
			gd.Error += 1
		}
		gd.bordSubData[getMouseX()][getMouseY()] = false
		if gd.selectedType == false {
			gd.selectedVal = getMouseX()
		} else {
			gd.selectedVal = getMouseY()
		}
		gd.selectedType = !gd.selectedType
	}
	if win {
		rl.DrawText("YOU WIN!!", 50, 50, 88, rl.RayWhite)
	}
	if lose {
		rl.DrawText("YOU LOSE!!", 50, 50, 88, rl.RayWhite)
	}
}
