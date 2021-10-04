package main

import (
	"math"
)

func generateGD() GameData {
	var gd GameData
	for i := 0; i < 11; i++ {
		for j := 0; j < 11; j++ {
			gd.Map[i][j] = false
		}
	}
	gd.Active = false
	gd.EnemysMoveVects[0] = up
	gd.EnemysMoveVects[1] = up
	gd.AiTimeBuffer = 0
	gd.poused = false
	//set map
	gd.Map[1][1] = true
	gd.Map[2][1] = true
	gd.Map[3][1] = true
	gd.Map[4][1] = true
	gd.Map[6][1] = true
	gd.Map[7][1] = true
	gd.Map[9][1] = true
	gd.Map[7][2] = true
	gd.Map[9][2] = true
	gd.Map[9][3] = true
	gd.Map[9][4] = true
	gd.Map[9][6] = true
	gd.Map[9][7] = true
	gd.Map[8][7] = true
	gd.Map[9][9] = true
	gd.Map[8][9] = true
	gd.Map[7][9] = true
	gd.Map[6][9] = true
	gd.Map[3][9] = true
	gd.Map[3][8] = true
	gd.Map[4][9] = true
	gd.Map[1][9] = true
	gd.Map[1][8] = true
	gd.Map[1][7] = true
	gd.Map[1][6] = true
	gd.Map[1][4] = true
	gd.Map[1][3] = true
	gd.Map[2][3] = true
	gd.Map[4][3] = true
	gd.Map[5][3] = true
	gd.Map[5][5] = true
	gd.Map[3][5] = true
	gd.Map[3][6] = true
	gd.Map[5][7] = true
	gd.Map[6][7] = true
	gd.Map[7][4] = true
	gd.Map[7][5] = true
	//
	gd.Enemys[0].y = 112
	gd.Enemys[0].x = 75 / 2
	gd.Enemys[1].y = 562
	gd.Enemys[1].x = 787
	gd.Player.x = 6*75 + 37
	gd.Player.y = 4*75 + 37
	gd.ItemsActive[0] = true
	gd.ItemsActive[1] = true
	gd.ItemsActive[2] = true
	gd.ItemsCollected = 0
	gd.Items[0].x = 4*75 + (75 / 2)
	gd.Items[0].y = 8*75 + (75 / 2)
	gd.Items[1].x = 9*75 + (75 / 2)
	gd.Items[1].y = 5*75 + (75 / 2)
	gd.Items[2].x = 3*75 + (75 / 2)
	gd.Items[2].y = 0*75 + (75 / 2)
	return gd
}

func checkIfPointInRect(x, y, RectX, RectY, SizeX, SizeY int32) bool {
	if x > RectX && y > RectY && x < RectX+SizeX && y < RectY+SizeY {
		return true
	}
	return false
}

func checkIfWin(gd GameData) bool {
	if gd.Player.y > 75*3 && gd.Player.y < 75*4 && gd.Player.x < 0 {
		return true
	}
	return false
}

func diference(val1, val2 int32) int32 {
	if val1 == val2 {
		return 0
	} else if val1 > val2 {
		return val2 - val1
	} else {
		return val1 - val2
	}
}

func calcDistance(x1, y1, x2, y2 int32) int32 {
	var xD = diference(x1, x2)
	var yD = diference(y1, y2)
	var buffer = (xD * xD) + (yD * yD)
	buffer = int32(math.Sqrt(float64(buffer)))
	return buffer
}

func checkIfLose(gd GameData) bool {
	for i := 0; i < 2; i++ {
		if calcDistance(gd.Enemys[i].x, gd.Enemys[i].y, gd.Player.x, gd.Player.y) < 50 {
			return true
		}
	}
	return false
}

func checkIfItemsPicked(gd GameData) int {
	for i := 0; i < 3; i++ {
		var x = gd.Items[i].x
		var y = gd.Items[i].y
		if gd.ItemsActive[i] {
			if calcDistance(x, y, gd.Player.x, gd.Player.y) < 14 {
				return i + 1
			}
		}
	}
	return 0
}
