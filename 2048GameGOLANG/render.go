package main

import rl "github.com/gen2brain/raylib-go/raylib"
import "strconv"

func render(gd GameData) {
	DrawBord(gd)
	DrawGrid()
	DrawScore(gd)
	rl.DrawRectangle(0, 900, 10000, 7, rl.Black)
	if gd.win {
		DrawWin(gd.score)
	}
}

func DrawWin(score int) {
	rl.DrawRectangle(0, 0, 1000, 1000, rl.Black)
	rl.DrawText("Score: "+strconv.Itoa(score), 15, 15, 40, rl.White)
	rl.DrawText("You Win Press Space To Reset", 50, 450, 50, rl.White)
}

func DrawScore(gd GameData) {
	rl.DrawRectangle(0, 900, 10000, 10000, rl.DarkGray)
	rl.DrawText("Score: "+strconv.Itoa(gd.score), 20, 917, 60, rl.RayWhite)
}

func DrawGrid() {
	for i := 1; i < 4; i++ {
		rl.DrawRectangle(0, int32(i*(900/4)), 1000, 7, rl.Black)
		rl.DrawRectangle(int32(i*(900/4)), 0, 7, 1000, rl.Black)
	}
}

func DrawBord(gd GameData) {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			var data = gd.bord[i][j]
			var color rl.Color = rl.RayWhite
			if data == 2 {
				color = Color_2
			}
			if data == 4 {
				color = Color_4
			}
			if data == 8 {
				color = Color_8
			}
			if data == 16 {
				color = Color_16
			}
			if data == 32 {
				color = Color_32
			}
			if data == 64 {
				color = Color_64
			}
			if data == 128 {
				color = Color_128
			}
			if data == 256 {
				color = Color_256
			}
			if data == 512 {
				color = Color_512
			}
			if data == 1024 {
				color = Color_1024
			}
			if data == 2048 {
				color = Color_2048
			}
			data_lenght := GetIntLenght(data)
			rl.DrawRectangle(int32(i*(900/4)), int32(j*(900/4)), (900 / 4), (900 / 4), color)
			rl.DrawText(strconv.Itoa(data), int32(i*(900/4)+115-(20*data_lenght)), int32(j*(900/4)+80), 70, rl.RayWhite)
		}
	}
}
