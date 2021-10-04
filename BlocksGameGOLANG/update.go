package main

import rl "github.com/gen2brain/raylib-go/raylib"

func update(gd *GameData) {
	if gd.state == 1 {
		if CheckIfClickedPice(ArrowPice, 28, 2) {
			gd.actives[0] = false
			gd.selectedIndex = 0
			gd.selectedVect = ArrowPice
			gd.state = 2
		}
		if CheckIfClickedPice(LinePice, 28, 5) {
			gd.actives[1] = false
			gd.selectedIndex = 1
			gd.selectedVect = LinePice
			gd.state = 2
		}
		if CheckIfClickedPice(SPice, 28, 8) {
			gd.actives[2] = false
			gd.selectedIndex = 2
			gd.selectedVect = SPice
			gd.state = 2
		}
		if CheckIfClickedPice(LPice, 28, 13) {
			gd.actives[3] = false
			gd.selectedIndex = 3
			gd.selectedVect = LPice
			gd.state = 2
		}
		if CheckIfClickedPice(BlockPice, 28, 16) {
			gd.actives[4] = false
			gd.selectedIndex = 4
			gd.selectedVect = BlockPice
			gd.state = 2
		}
	} else if gd.state == 2 {
		for i := 0; i < 4; i++ {
			rl.DrawRectangle(int32(gd.selectedVect[i].x)*50+int32(rl.GetMouseX())-25, int32(gd.selectedVect[i].y)*50-25+int32(rl.GetMouseY()), 50, 50, rl.Black)
			rl.DrawRectangle(int32(gd.selectedVect[i].x)*50+int32(rl.GetMouseX())+1-25, int32(gd.selectedVect[i].y)*50-25+int32(rl.GetMouseY())+1, 48, 48, rl.Blue)
		}
		if !rl.IsMouseButtonDown(rl.MouseLeftButton) {
			gd.actives[gd.selectedIndex] = true
			gd.state = 1
			if CheckIfCanBeInBord(*gd) {
				PutPice(gd)
				gd.score++
			}
		}
	}
}
