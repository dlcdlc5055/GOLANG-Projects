package main

import rl "github.com/gen2brain/raylib-go/raylib"

func render(gd GameData) {
	renderBg()
	renderBord(gd)
	renderGrid()
	renderSunks(gd)
}

func renderBg() {
	rl.DrawRectangle(10, 10, 620, 620, rl.Black)
	rl.DrawRectangle(1600-20-610, 10, 620, 620, rl.Black)
	rl.DrawRectangle(20, 20, 600, 600, rl.Gray)
	rl.DrawRectangle(1600-20-600, 20, 600, 600, rl.Gray)
}

func renderGrid() {
	for i := 0; i < 11; i++ {
		//bord 1
		rl.DrawRectangle(int32(20+i*60), 20, 1, 600, rl.RayWhite)
		rl.DrawRectangle(20, int32(20+i*60), 600, 1, rl.RayWhite)
		//bord 2
		rl.DrawRectangle(1600-int32(20+i*60), 20, 1, 600, rl.RayWhite)
		rl.DrawRectangle(1600-20-600, int32(20+i*60), 600, 1, rl.RayWhite)
	}
}

func renderSunks(gd GameData) {
	for i := 0; i < 5; i++ {
		if gd.sunkenShipsPlayerRed[i] {
			rl.DrawCircle(int32(100+i*100), 660, 12, rl.Red)
		} else {
			rl.DrawCircle(int32(100+i*100), 660, 12, rl.Gray)
			rl.DrawCircle(int32(100+i*100), 660, 8, rl.RayWhite)
		}
		if gd.sunkenShipsPlayerBlue[i] {
			rl.DrawCircle(1575-int32(100+i*100), 660, 12, rl.Blue)
		} else {
			rl.DrawCircle(1575-int32(100+i*100), 660, 12, rl.Gray)
			rl.DrawCircle(1575-int32(100+i*100), 660, 8, rl.RayWhite)
		}
	}
}

func renderBord(gd GameData) {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if gd.bordBlue[i][j] == ship || gd.bordBlue[i][j] == hitShip {
				rl.DrawRectangle(20+int32(60*i), 20+int32(60*j), 60, 60, rl.Blue)
			}
			if gd.bordRed[i][j] == ship || gd.bordRed[i][j] == hitShip {
				rl.DrawRectangle(1600-20-int32(60*i)-600, 20+int32(60*j), 60, 60, rl.Red)
			}
			if gd.bordRed[i][j] == hitShip {
				rl.DrawCircle(1600-20-int32(60*i)+30-600, 20+int32(60*j)+30, 10, rl.Pink)
			}
			if gd.bordBlue[i][j] == hitShip {
				rl.DrawCircle(20+int32(60*i)+30, 20+int32(60*j)+30, 10, rl.Pink)
			}
		}
	}
}
