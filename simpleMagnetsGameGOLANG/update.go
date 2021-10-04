package main

import rl "github.com/gen2brain/raylib-go/raylib"

func update(gd *GameData) {
	if rl.IsKeyPressed(rl.KeyEnter) {
		*gd = generateGameData()
	}

	var win = checkIfWin(*gd)
	var lose = false

	if gd.nrMoves > 25 {
		lose = true
	}

	if win {
		rl.DrawRectangle(0, 0, 100000, 100000, rl.Black)
		rl.DrawText("YOU WIN!!", 100, 100, 45, rl.RayWhite)
		rl.DrawText("PRESS ENTER TO RESET!!", 100, 400, 45, rl.RayWhite)
	}

	if lose {
		rl.DrawRectangle(0, 0, 100000, 100000, rl.Black)
		rl.DrawText("YOU LOSE!!", 100, 100, 45, rl.RayWhite)
		rl.DrawText("PRESS ENTER TO RESET!!", 100, 400, 45, rl.RayWhite)
	}

	var clicked = getClickedBtn()

	if !win && !lose {
		for i := 0; i < 7; i++ {
			if i == clicked {
				gd.nrMoves += 1
				gd.player[i] = !gd.player[i]
				gd.player[i+1] = !gd.player[i+1]
			}
		}
	}
}
