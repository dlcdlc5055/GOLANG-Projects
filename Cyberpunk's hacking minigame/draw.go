package main

import (
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func draw(gd GameData) {
	var win = checkForWin(gd)
	var lose = checkForLose(gd)
	drawBG(gd)
	drawBord(gd)
	drawCorrects(gd)
	if win || lose {
		rl.DrawRectangle(0, 0, 1000, 1000, rl.Black)
	}
}

func drawBG(gd GameData) {
	rl.DrawRectangle(0, 0, 10000, 10000, rl.DarkGray)
	rl.DrawRectangle(50, 50, 85*5, 85*5, rl.LightGray)
	rl.DrawRectangle(100+85*5, 50, 85*3, 65, rl.LightGray)
	rl.DrawRectangle(100+85*5, 50, 85*int32(gd.corrects), 65, rl.DarkGreen)
	rl.DrawRectangle(100+85*5, 200, 85*3, 65, rl.LightGray)
	rl.DrawRectangle(100+85*5, 200, 85*int32(gd.Error), 65, rl.Red)
	rl.DrawText("errors: "+strconv.Itoa(gd.Error), 100+85*5+15, 210, 45, rl.Black)
	if gd.selectedType == false {
		rl.DrawRectangle(50, 50+85*int32(gd.selectedVal), 85*5, 85, rl.DarkGreen)
		if getMouseY() == gd.selectedVal && getMouseX() > -1 {
			rl.DrawRectangle(50+85*int32(getMouseX()), 50+85*int32(gd.selectedVal), 85, 85, rl.Red)
		}
	} else {
		rl.DrawRectangle(50+85*int32(gd.selectedVal), 50, 85, 85*5, rl.DarkGreen)
		if getMouseX() == gd.selectedVal && getMouseY() > -1 {
			rl.DrawRectangle(50+85*int32(gd.selectedVal), 50+85*int32(getMouseY()), 85, 85, rl.Red)
		}
	}
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if gd.bordSubData[i][j] == false {
				rl.DrawRectangle(50+85*int32(i), 50+85*int32(j), 85, 85, rl.LightGray)
			}
		}
	}
}

func drawBord(gd GameData) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			var buffer = gd.bord[i][j]
			if gd.bordSubData[i][j] {
				rl.DrawText(buffer, 50+int32(i)*85+9, 50+int32(j)*85+18, 55, rl.Black)
			} else {
				rl.DrawText(buffer, 50+int32(i)*85+9, 50+int32(j)*85+18, 55, rl.DarkGray)
			}
		}
	}
}

func drawCorrects(gd GameData) {
	const base_x = 100 + 85*5
	for i := 0; i < 3; i++ {
		var buffer = gd.correct[i]
		rl.DrawText(buffer, base_x+int32(i)*85+5, 55, 55, rl.Black)
	}
}
