package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func draw(gd *GameData) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if (i%2 == 1 && j%2 == 0) || (i%2 == 0 && j%2 == 1) {
				rl.DrawRectangle(int32(i)*79, int32(j)*79, 79, 79, rl.DarkBlue)
			} else {
				rl.DrawRectangle(int32(i)*79, int32(j)*79, 79, 79, rl.SkyBlue)
			}
		}
	}
	drawPolice(*gd)
	drawFinish(*gd)
	drawPlayer(*gd)
	drawMap(*gd)
}

func drawMap(gd GameData) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			var block = gd.Map[i][j]
			if block.up {
				rl.DrawRectangle(int32(i)*79, int32(j)*79-2, 79, 5, rl.Black)
			}
			if block.down {
				rl.DrawRectangle(int32(i)*79, int32(j)*79-2+79, 79, 5, rl.Black)
			}
			if block.left {
				rl.DrawRectangle(int32(i)*79-2, int32(j)*79, 5, 79, rl.Black)
			}
			if block.right {
				rl.DrawRectangle(int32(i)*79-2+79, int32(j)*79, 5, 79, rl.Black)
			}
		}
	}
}

func drawPlayer(gd GameData) {
	var x = gd.Player.X
	var y = gd.Player.Y
	rl.DrawCircle(x*79+(79/2), y*79+(79/2), 33, rl.Green)
}

func drawPolice(gd GameData) {
	//drawPBuffer(gd)
	var x = gd.Police.X
	var y = gd.Police.Y
	rl.DrawCircle(x*79+(79/2), y*79+(79/2), 33, rl.Blue)
}

func drawFinish(gd GameData) {
	var x = gd.Finish.X
	var y = gd.Finish.Y
	rl.DrawRectangle(x*79, y*79, 79, 79, rl.Red)
}

func drawWin() {
	rl.DrawRectangle(0, 0, 10000, 10000, rl.Black)
	rl.DrawText("You Win!!", 50, 50, 45, rl.RayWhite)
	rl.DrawText("Press Space To Reset!!", 50, 150, 45, rl.RayWhite)
}

func drawLose() {
	rl.DrawRectangle(0, 0, 10000, 10000, rl.Black)
	rl.DrawText("You Lose!!", 50, 50, 45, rl.RayWhite)
	rl.DrawText("Press Space To Reset!!", 50, 150, 45, rl.RayWhite)
}

func drawPBuffer(gd GameData) {
	if gd.PoliceMoved {
		var x = gd.PoliceBUffer.X
		var y = gd.PoliceBUffer.Y
		rl.DrawCircle(x*79+(79/2), y*79+(79/2), 33, rl.DarkPurple)
	}
}
