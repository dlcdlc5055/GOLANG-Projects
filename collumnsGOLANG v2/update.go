package main

import (
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const typesOfPices int = 6

func update(gd *GameData) {

	if rl.IsKeyPressed(rl.KeyM) {
		gd.mode = !gd.mode
	}

	if gd.pice.pos.y == 10 {
		PlacePice(gd)
	}

	gd.timeBuffer += rl.GetFrameTime()

	var delay float32 = .6

	if gd.timeBuffer > delay && !checkIfOcupied(*gd, gd.pice.pos, 0, 1) {
		gd.timeBuffer = 0
		gd.pice.pos.y += 1
	} else if checkIfOcupied(*gd, gd.pice.pos, 0, 1) {
		PlacePice(gd)
	}

	if rl.IsKeyPressed(rl.KeyA) && !checkIfOcupied(*gd, gd.pice.pos, -1, 0) {
		gd.pice.pos.x -= 1
		if gd.pice.pos.x < 0 {
			gd.pice.pos.x = 0
		}
	}

	if rl.IsKeyPressed(rl.KeyD) && !checkIfOcupied(*gd, gd.pice.pos, 1, 0) {
		gd.pice.pos.x += 1
		if gd.pice.pos.x > 5 {
			gd.pice.pos.x = 5
		}
	}

	if rl.IsKeyPressed(rl.KeyW) {
		gd.pice = turnPice(gd.pice)
	}

	if rl.IsKeyPressed(rl.KeyS) {
		gd.pice = reverseTurnPice(gd.pice)
	}

	if rl.IsKeyPressed(rl.KeyE) {
		gd.nextPice = generatePice(0, gd)
	}

	if rl.IsKeyPressed(rl.KeyQ) {
		gd.nextPice = generatePice(0, gd)
	}

	var b1 = false

	if rl.IsKeyPressed(rl.KeySpace) {
		if gd.bord[gd.pice.pos.x][12] < 0 {
			gd.pice.pos.y = int32(12) - 2
			PlacePice(gd)
			b1 = true
		} else {
			for i := 0; i < 13; i++ {
				if gd.bord[gd.pice.pos.x][i] > -1 && b1 == false {
					gd.pice.pos.y = int32(i) - 3
					PlacePice(gd)
					b1 = true
				}
			}
		}
	}

	processFall(gd)
	ProcessMaches(gd)
	processFall(gd)

	var full = checkIfBordFull(*gd)

	if rl.IsKeyPressed(rl.KeyEnter) {
		*gd = generateGameData()
	}

	if full {
		rl.DrawRectangle(0, 0, 10000, 10000, rl.Black)
		rl.DrawText("Final Score: "+strconv.Itoa(gd.score), 50, 50, 50, rl.RayWhite)
		rl.DrawText("Press Enter To Reset", 50, 200, 50, rl.RayWhite)
	}

	if gd.mode {
		generatePossiblesForEasyMode(gd)
	}
}
