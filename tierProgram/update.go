package main

import rl "github.com/gen2brain/raylib-go/raylib"

func update(circles *[]Circle, selectedCircleIndex *int, isSelectedCircle *bool, vectOffset *Vector, DoneCirclesCount *[7]int8, DoneCirclesColors *[7][7]rl.Color, circlesBackUp *[]Circle) {
	if rl.IsKeyReleased(rl.KeySpace) && CheckForWin(*circles) {
		reset(circles, selectedCircleIndex, isSelectedCircle, vectOffset, DoneCirclesCount, DoneCirclesColors, circlesBackUp)
	}
	//var colors = []rl.Color{rl.Red, rl.Orange, rl.Green, rl.Yellow, rl.Blue, rl.Purple, rl.Pink}
	var Mpos = rl.GetMousePosition()
	if rl.IsMouseButtonPressed(rl.MouseLeftButton) && !*isSelectedCircle {
		var index = getClickedCircleIndex(*circles, int32(Mpos.X), int32(Mpos.Y))
		vectOffset.y = 0
		vectOffset.x = 0
		if index != -1 && !(*circles)[index].done && Mpos.X > 900 {
			(*circles)[index].active = false
			*isSelectedCircle = true
			*selectedCircleIndex = index
			getOffSetForCircle(Mpos.X, Mpos.Y, (*circles)[index].x, (*circles)[index].y, vectOffset)
		} else if Mpos.X < 900 && Mpos.X > 128 {
			for i := 0; i < 7; i++ {
				for j := 0; j < int(DoneCirclesCount[i]); j++ {
					var b1 = int32(138 + 50 + 10 + j*108)
					var b2 = int32(i*128 + 50 + 12)
					var buffer = CirclePointCollision(b1, b2, int32(Mpos.X), int32(Mpos.Y), 54, false)
					if buffer {
						var index = getClickedCircleIndex(*circles, int32(Mpos.X), int32(Mpos.Y))
						vectOffset.y = 0
						vectOffset.x = 0
						if index != -1 {
							(*circles)[index].active = false
							*isSelectedCircle = true
							*selectedCircleIndex = index
							getOffSetForCircle(Mpos.X, Mpos.Y, (*circles)[index].x, (*circles)[index].y, vectOffset)
							RemoveCircle(j, i, DoneCirclesCount, DoneCirclesColors, circles)
						}
					}
				}
			}
		}
	} else if rl.IsMouseButtonReleased(rl.MouseLeftButton) {
		var id = int(Mpos.Y) / 128
		if id < 0 {
			id = 0
		} else if id > 6 {
			id = 6
		}
		if Mpos.X < 900 && *isSelectedCircle && (*DoneCirclesCount)[id] < 7 {

			(*DoneCirclesColors)[(*DoneCirclesCount)[id]][id] = (*circles)[*selectedCircleIndex].color
			(*DoneCirclesCount)[id] += 1
			(*circles)[*selectedCircleIndex].active = true
			(*circles)[*selectedCircleIndex].done = true
			var b1 = int32(90 + int32((*DoneCirclesCount)[id])*108)
			var b2 = int32(id*128 + 50 + 12)
			(*circles)[*selectedCircleIndex].x = b1
			(*circles)[*selectedCircleIndex].y = b2
			//rest
			*selectedCircleIndex = -1
			*isSelectedCircle = false

		} else if *selectedCircleIndex != -1 {
			//restore
			(*circles)[*selectedCircleIndex].active = true
			(*circles)[*selectedCircleIndex].done = false
			//rest
			(*circles)[*selectedCircleIndex].x = (*circlesBackUp)[*selectedCircleIndex].x
			(*circles)[*selectedCircleIndex].y = (*circlesBackUp)[*selectedCircleIndex].y
			*isSelectedCircle = false
			*selectedCircleIndex = -1
		}
	} else if rl.IsMouseButtonDown(rl.MouseLeftButton) && *isSelectedCircle {
		var Mpos = rl.GetMousePosition()
		rl.DrawCircle(int32(Mpos.X)+vectOffset.x, int32(Mpos.Y)+vectOffset.y, float32(50+4), rl.Black)
		rl.DrawCircle(int32(Mpos.X)+vectOffset.x, int32(Mpos.Y)+vectOffset.y, 50, (*circles)[*selectedCircleIndex].color)
	}
	if CheckForWin(*circles) {
		drawWin()
	}
}
