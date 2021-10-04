package main

import (
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func draw(gd GameData) {
	drawGrid()
	drawMap(gd)
	drawItems(gd)
	drawPlayer(gd)
	drawEnemys(gd)
}

func drawMap(gd GameData) {
	const xM = -80
	const yM = -80
	for i := 0; i < 11; i++ {
		for j := 0; j < 10; j++ {
			if gd.Map[i][j] == Block {
				rl.DrawRectangle(int32(i)*80+1+xM, int32(j)*80+1+yM, 80, 80, rl.Blue)
			} else if gd.Map[i][j] == Win {
				rl.DrawRectangle(int32(i)*80+2+xM, int32(j)*80+2+yM, 80-1, 80-1, rl.Red)
			}
		}
	}
}

func drawGrid() {
	for i := 0; i < 11; i++ {
		rl.DrawRectangle(int32(i)*80, 0, 2, 1000, rl.Black)
	}
	for i := 0; i < 10; i++ {
		rl.DrawRectangle(0, int32(i)*80, 1000, 2, rl.Black)
	}
}

func drawEnemys(gd GameData) {
	for i := 0; i < 2; i++ {
		var x = gd.Enemys[i].x
		var y = gd.Enemys[i].y
		rl.DrawCircle(x*80+40, y*80+40, 35, rl.Red)
	}
}

func drawItems(gd GameData) {
	for i := 0; i < 3; i++ {
		var x = gd.Items[i].x
		var y = gd.Items[i].y
		if gd.ItemsActive[i] {
			rl.DrawCircle(x*80+40, y*80+40, 35, rl.Orange)
		}
	}
}

func drawPlayer(gd GameData) {
	var x = gd.Player.x
	var y = gd.Player.y
	rl.DrawCircle(x*80+40, y*80+40, 35, rl.Green)
}

func drawWin(gd GameData) {
	rl.DrawRectangle(0, 0, 1000, 1000, rl.Black)
	rl.DrawText("YOU WIN "+strconv.Itoa(gd.ColectedItems)+" ITEMS!!", 50, 50, 50, rl.RayWhite)
}

func drawLose() {
	rl.DrawRectangle(0, 0, 1000, 1000, rl.Black)
	rl.DrawText("YOU LOSE!!", 50, 50, 50, rl.RayWhite)
}
