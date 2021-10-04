package main

import (
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func draw(gd GameData) {
	rl.DrawRectangle(0, 3*75, 75, 75, rl.Red)
	drawGrid()
	drawBackground(gd)
	rl.DrawRectangle(0, 0-20, 75, 75*3-20, rl.RayWhite)
	rl.DrawRectangle(0, 75*4+4, 75, 75*500, rl.RayWhite)
	drawItemes(gd)
	drawPlayer(gd)
	//	drawEnemys(gd)
	rl.DrawRectangle(0, 825, 1000, 1000, rl.Black)
	rl.DrawText("Items Collected : "+strconv.Itoa(gd.ItemsCollected), 20, 835, 27, rl.RayWhite)
	if gd.Win == true {
		drawWin(gd)
	} else if gd.Lose == true {
		drawLose(gd)
	}
}

func drawWin(gd GameData) {
	rl.DrawRectangle(0, 0, 1000, 1999, rl.Black)
	if gd.ItemsCollected < 3 {
		rl.DrawText("You Collected "+strconv.Itoa(gd.ItemsCollected)+" Items!!", 50, 50, 70, rl.RayWhite)
	} else {
		rl.DrawText("You Collected All Items!!", 20, 50, 70, rl.RayWhite)
	}
}

func drawLose(gd GameData) {
	rl.DrawRectangle(0, 0, 1000, 1999, rl.Black)
	rl.DrawText("You Lose!!", 50, 50, 70, rl.RayWhite)
}

func drawItemes(gd GameData) {
	for i := 0; i < 3; i++ {
		if gd.ItemsActive[i] {
			var x = gd.Items[i].x
			var y = gd.Items[i].y
			rl.DrawCircle(x+75, y, 35, rl.SkyBlue)
		}
	}
}

func drawBackground(gd GameData) {
	for i := 0; i < 11; i++ {
		for j := 0; j < 11; j++ {
			var value = gd.Map[i][j]
			if value {
				rl.DrawRectangle(int32(i)*75+75, int32(j)*75, 75, 75, rl.Blue)
			}
		}
	}
}

func drawGrid() {
	for i := 0; i < 11; i++ {
		rl.DrawRectangle(0, int32(i)*75, 1000, 2, rl.Black)
		rl.DrawRectangle(int32(i)*75+75, 0, 2, 1000, rl.Black)
	}
}

func drawEnemys(gd GameData) {
	rl.DrawCircle(gd.Enemys[0].x+75, gd.Enemys[0].y, 35, rl.Red)
	rl.DrawCircle(gd.Enemys[1].x+75, gd.Enemys[1].y, 35, rl.Red)
}

func drawPlayer(gd GameData) {
	rl.DrawCircle(gd.Player.x+75, gd.Player.y, 35, rl.Green)
}
