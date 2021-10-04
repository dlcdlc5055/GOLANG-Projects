package main

import (
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func render(gd GameData) {
	DrawBord(gd)
	DrawGrid()
	DrawSidePices(gd)
	DrawScore(gd.score)
}

func DrawBord(gd GameData) {
	for i := 0; i < 25; i++ {
		for j := 0; j < 18; j++ {
			if gd.bord[i][j] {
				rl.DrawRectangle(int32(i*50), int32(j*50), 50, 50, rl.Blue)
			}
		}
	}
}

func DrawGrid() {
	for i := 1; i < 26; i++ {
		rl.DrawRectangle(int32(i*50), 0, 1, 1000, rl.Black)
	}
	for i := 1; i < 19; i++ {
		rl.DrawRectangle(0-50, int32(i*50), 1300, 1, rl.Black)
	}
}

func DrawScore(score int) {
	rl.DrawText("Score: "+strconv.Itoa(score), 1265, 15, 25, rl.Black)
}

func DrawSidePices(gd GameData) {
	for i := 0; i < 4; i++ {
		if gd.actives[0] {
			rl.DrawRectangle(int32((28+ArrowPice[i].x)*50), int32((2+ArrowPice[i].y)*50), 50, 50, rl.Black)
			rl.DrawRectangle(int32((28+ArrowPice[i].x)*50)+1, int32((2+ArrowPice[i].y)*50)+1, 48, 48, rl.Blue)
		}
		if gd.actives[1] {
			rl.DrawRectangle(int32((28+LinePice[i].x)*50), int32((5+LinePice[i].y)*50), 50, 50, rl.Black)
			rl.DrawRectangle(int32((28+LinePice[i].x)*50)+1, int32((5+LinePice[i].y)*50)+1, 48, 48, rl.Blue)
		}
		if gd.actives[2] {
			rl.DrawRectangle(int32((28+SPice[i].x)*50), int32((8+SPice[i].y)*50), 50, 50, rl.Black)
			rl.DrawRectangle(int32((28+SPice[i].x)*50)+1, int32((8+SPice[i].y)*50)+1, 48, 48, rl.Blue)
		}
		if gd.actives[3] {
			rl.DrawRectangle(int32((28+LPice[i].x)*50), int32((13+LPice[i].y)*50), 50, 50, rl.Black)
			rl.DrawRectangle(int32((28+LPice[i].x)*50)+1, int32((13+LPice[i].y)*50)+1, 48, 48, rl.Blue)
		}
		if gd.actives[4] {
			rl.DrawRectangle(int32((28+BlockPice[i].x)*50), int32((16+BlockPice[i].y)*50), 50, 50, rl.Black)
			rl.DrawRectangle(int32((28+BlockPice[i].x)*50)+1, int32((16+BlockPice[i].y)*50)+1, 48, 48, rl.Blue)
		}
	}
}
