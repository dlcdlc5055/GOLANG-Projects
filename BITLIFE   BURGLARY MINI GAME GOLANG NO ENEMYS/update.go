package main

import (
	"math/rand"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func update(gd *GameData) {
	//var lose = checkIfLose(*gd)
	var win = checkIfWin(*gd)
	gd.Win = win
	//	gd.Lose = lose
	if gd.Win {
		gd.poused = true
	}
	if checkIfItemsPicked(*gd) != 0 {
		gd.ItemsActive[checkIfItemsPicked(*gd)-1] = false
		gd.ItemsCollected += 1
	}
	if gd.PlayerMoveVect != 0 {
		if rl.IsKeyDown(rl.KeyW) && gd.PlayerMoveVect != down && checkIfPossibleUpLong(*gd) && checkIfCloseToCenter(*gd) && gd.Player.y > 75 {
			gd.PlayerMoveVect = up
		}
		if rl.IsKeyDown(rl.KeyS) && gd.PlayerMoveVect != up && checkIfPossibleDownLong(*gd) && checkIfCloseToCenter(*gd) && gd.Player.y < 75*10 {
			gd.PlayerMoveVect = down
		}
		if rl.IsKeyDown(rl.KeyD) && gd.PlayerMoveVect != left && checkIfPossibleRightLong(*gd) && checkIfCloseToCenter(*gd) && gd.Player.x < 75*10 {
			gd.PlayerMoveVect = right
		}
		if rl.IsKeyDown(rl.KeyA) && gd.PlayerMoveVect != right && checkIfPossibleLeftLong(*gd) && checkIfCloseToCenter(*gd) && gd.Player.x > 75 {
			gd.PlayerMoveVect = left
		}
		if rl.IsKeyDown(rl.KeyA) && gd.PlayerMoveVect != right && gd.Player.x < 75 && gd.Player.y > 75*3 && gd.Player.y < 75*4 && checkIfCloseToCenter(*gd) {
			gd.PlayerMoveVect = left
		}

	} else {
		if rl.IsKeyDown(rl.KeyW) && checkIfPossibleUpLong(*gd) && checkIfCloseToCenter(*gd) {
			gd.Active = true
			gd.PlayerMoveVect = up
		}
		if rl.IsKeyDown(rl.KeyS) && checkIfPossibleDownLong(*gd) && checkIfCloseToCenter(*gd) {
			gd.Active = true
			gd.PlayerMoveVect = down
		}
		if rl.IsKeyDown(rl.KeyD) && checkIfPossibleRightLong(*gd) && checkIfCloseToCenter(*gd) {
			gd.Active = true
			gd.PlayerMoveVect = right
		}
		if rl.IsKeyDown(rl.KeyA) && checkIfPossibleLeftLong(*gd) && checkIfCloseToCenter(*gd) {
			gd.Active = true
			gd.PlayerMoveVect = left
		}
	}
	updatePlayer(gd, 6)
	if gd.Active {
		//	updateEnemys(gd, 6)
	}
	//	ProcessAI(gd)
}

func updateEnemys(gd *GameData, speed int32) {
	for i := 0; i < 2; i++ {
		if gd.EnemysMoveVects[i] == up && checkIfPossibleUpAI(*gd, i) && gd.Enemys[i].y > 40 {
			gd.Enemys[i].y -= speed
		}
		if gd.EnemysMoveVects[i] == down && checkIfPossibleDownAI(*gd, i) && gd.Enemys[i].y < 825-50 {
			gd.Enemys[i].y += speed
		}
		if gd.EnemysMoveVects[i] == left && checkIfPossibleLeftAI(*gd, i) && gd.Enemys[i].x > 40 {
			gd.Enemys[i].x -= speed
		}
		if gd.EnemysMoveVects[i] == right && checkIfPossibleRightAI(*gd, i) && gd.Enemys[i].x < 825-40 {
			gd.Enemys[i].y += speed
		}
	}
}

func ProcessAI(gd *GameData) {
	if gd.AiTimeBuffer > 1.1 {
		gd.AiTimeBuffer = 0
		for i := 0; i < 2; i++ {
			rand.Seed(time.Now().UnixNano())
			var dirBuffer = rand.Intn(4) + 1
			for true {
				if !checkIfCloseToCenterAI(*gd, i) {
					break
				}
				if !checkIfPossibleAllAI(*gd, i) {
					break
				}
				if dirBuffer == up && checkIfPossibleUpLongAI(*gd, i) && gd.EnemysMoveVects[i] != down && checkIfCloseToCenterAI(*gd, i) {
					gd.EnemysMoveVects[i] = dirBuffer
					break
				} else if dirBuffer == down && checkIfPossibleDownLongAI(*gd, i) && gd.EnemysMoveVects[i] != up && checkIfCloseToCenterAI(*gd, i) {
					gd.EnemysMoveVects[i] = dirBuffer
					break
				} else if dirBuffer == left && checkIfPossibleLeftLongAI(*gd, i) && gd.EnemysMoveVects[i] != right && checkIfCloseToCenterAI(*gd, i) {
					gd.EnemysMoveVects[i] = dirBuffer
					break
				} else if dirBuffer == right && checkIfPossibleRightLongAI(*gd, i) && gd.EnemysMoveVects[i] != left && checkIfCloseToCenterAI(*gd, i) {
					gd.EnemysMoveVects[i] = dirBuffer
					break
				} else {
					dirBuffer = rand.Intn(4) + 1
				}
			}
		}
	}
	gd.AiTimeBuffer += rl.GetFrameTime()
}

func updatePlayer(gd *GameData, speed int32) {
	if gd.PlayerMoveVect == up && checkIfPossibleUp(*gd) && gd.Player.y > 40 {
		gd.Player.y -= speed
	}
	if gd.PlayerMoveVect == down && checkIfPossibleDown(*gd) && gd.Player.y < 825-40 {
		gd.Player.y += speed
	}
	if gd.PlayerMoveVect == left && checkIfPossibleLeft(*gd) && gd.Player.x > 40 {
		gd.Player.x -= speed
	}
	if gd.PlayerMoveVect == right && checkIfPossibleRight(*gd) && gd.Player.x < 825-40 {
		gd.Player.x += speed
	}
	if gd.PlayerMoveVect == left && gd.Player.y > 75*3 && gd.Player.y < 75*4 && gd.Player.x < 75 {
		gd.Player.x -= speed
	}
}

func checkIfCloseToCenter(gd GameData) bool {
	for i := 0; i < 11; i++ {
		for j := 0; j < 11; j++ {
			var xCircle = int32(i*75 + (75 / 2))
			var yCircle = int32(j*75 + (75 / 2))
			if calcDistance(xCircle, yCircle, gd.Player.x, gd.Player.y) < 13 {
				return true
			}
		}
	}
	return false
}

func checkIfPossibleAll(gd GameData) bool {
	var v1 = checkIfPossibleLeft(gd)
	var v2 = checkIfPossibleRight(gd)
	var v3 = checkIfPossibleDown(gd)
	var v4 = checkIfPossibleUp(gd)
	if v1 == false && v2 == false && v3 == false && v4 == false {
		return true
	}
	return false
}

//possible AI\

func checkIfCloseToCenterAI(gd GameData, index int) bool {
	for i := 0; i < 11; i++ {
		for j := 0; j < 11; j++ {
			var xCircle = int32(i*75 + (75 / 2))
			var yCircle = int32(j*75 + (75 / 2))
			if calcDistance(xCircle, yCircle, gd.Enemys[index].x, gd.Enemys[index].y) < 13 {
				return true
			}
		}
	}
	return false
}

func checkIfPossibleUpAI(gd GameData, index int) bool {
	var x = gd.Enemys[index].x
	var y = gd.Enemys[index].y - 37
	for i := 0; i < 11; i++ {
		for j := 0; j < 11; j++ {
			if gd.Map[i][j] {
				if checkIfPointInRect(x, y, int32(i)*75, int32(j)*75, 75, 75) && y > 75 {
					return false
				}
			}
		}
	}
	return true
}

func checkIfPossibleDownAI(gd GameData, index int) bool {
	var x = gd.Enemys[index].x
	var y = gd.Enemys[index].y + 37
	for i := 0; i < 11; i++ {
		for j := 0; j < 11; j++ {
			if gd.Map[i][j] {
				if checkIfPointInRect(x, y, int32(i)*75, int32(j)*75, 75, 75) && y < 75*10 {
					return false
				}
			}
		}
	}
	return true
}

func checkIfPossibleLeftAI(gd GameData, index int) bool {
	var x = gd.Enemys[index].x - 37
	var y = gd.Enemys[index].y
	for i := 0; i < 11; i++ {
		for j := 0; j < 11; j++ {
			if gd.Map[i][j] {
				if checkIfPointInRect(x, y, int32(i)*75, int32(j)*75, 75, 75) && x > 75 {
					return false
				}
			}
		}
	}
	return true
}

func checkIfPossibleRightAI(gd GameData, index int) bool {
	var x = gd.Enemys[index].x + 37
	var y = gd.Enemys[index].y
	for i := 0; i < 11; i++ {
		for j := 0; j < 11; j++ {
			if gd.Map[i][j] {
				if checkIfPointInRect(x, y, int32(i)*75, int32(j)*75, 75, 75) && x < 75*10 {
					return false
				}
			}
		}
	}
	return true
}

//AI Long
func checkIfPossibleUpLongAI(gd GameData, index int) bool {
	var x = gd.Enemys[index].x
	var y = gd.Enemys[index].y - 85
	for i := 0; i < 11; i++ {
		for j := 0; j < 11; j++ {
			if gd.Map[i][j] {
				if checkIfPointInRect(x, y, int32(i)*75, int32(j)*75, 75, 75) || checkIfPointInRect(x+30, y, int32(i)*75, int32(j)*75, 75, 75) || checkIfPointInRect(x-30, y, int32(i)*75, int32(j)*75, 75, 75) && y > 75 {
					return false
				}
			}
		}
	}
	return true
}

func checkIfPossibleDownLongAI(gd GameData, index int) bool {
	var x = gd.Enemys[index].x
	var y = gd.Enemys[index].y + 85
	for i := 0; i < 11; i++ {
		for j := 0; j < 11; j++ {
			if gd.Map[i][j] {
				if checkIfPointInRect(x, y, int32(i)*75, int32(j)*75, 75, 75) || checkIfPointInRect(x+30, y, int32(i)*75, int32(j)*75, 75, 75) || checkIfPointInRect(x-30, y, int32(i)*75, int32(j)*75, 75, 75) && y < 75*10 {
					return false
				}
			}
		}
	}
	return true
}

func checkIfPossibleLeftLongAI(gd GameData, index int) bool {
	var x = gd.Enemys[index].x - 85
	var y = gd.Enemys[index].y
	for i := 0; i < 11; i++ {
		for j := 0; j < 11; j++ {
			if gd.Map[i][j] {
				if checkIfPointInRect(x, y, int32(i)*75, int32(j)*75, 75, 75) || checkIfPointInRect(x, y+30, int32(i)*75, int32(j)*75, 75, 75) || checkIfPointInRect(x, y-30, int32(i)*75, int32(j)*75, 75, 75) && x > 75 {
					return false
				}
			}
		}
	}
	return true
}

func checkIfPossibleRightLongAI(gd GameData, index int) bool {
	var x = gd.Enemys[index].x + 85
	var y = gd.Enemys[index].y
	for i := 0; i < 11; i++ {
		for j := 0; j < 11; j++ {
			if gd.Map[i][j] {
				if checkIfPointInRect(x, y, int32(i)*75, int32(j)*75, 75, 75) || checkIfPointInRect(x, y+30, int32(i)*75, int32(j)*75, 75, 75) || checkIfPointInRect(x, y-30, int32(i)*75, int32(j)*75, 75, 75) && x < 75*10 {
					return false
				}
			}
		}
	}
	return true
}

func checkIfPossibleAllAI(gd GameData, i int) bool {
	var v1 = checkIfPossibleLeftLongAI(gd, i)
	var v2 = checkIfPossibleRightLongAI(gd, i)
	var v3 = checkIfPossibleDownLongAI(gd, i)
	var v4 = checkIfPossibleUpLongAI(gd, i)
	if v1 == true || v2 == true || v3 == true || v4 == true {
		return true
	}
	return false
}

// possible player

func checkIfPossibleUp(gd GameData) bool {
	var x = gd.Player.x
	var y = gd.Player.y - 37
	for i := 0; i < 11; i++ {
		for j := 0; j < 11; j++ {
			if gd.Map[i][j] {
				if checkIfPointInRect(x, y, int32(i)*75, int32(j)*75, 75, 75) {
					return false
				}
			}
		}
	}
	return true
}

func checkIfPossibleDown(gd GameData) bool {
	var x = gd.Player.x
	var y = gd.Player.y + 37
	for i := 0; i < 11; i++ {
		for j := 0; j < 11; j++ {
			if gd.Map[i][j] {
				if checkIfPointInRect(x, y, int32(i)*75, int32(j)*75, 75, 75) {
					return false
				}
			}
		}
	}
	return true
}

func checkIfPossibleLeft(gd GameData) bool {
	var x = gd.Player.x - 37
	var y = gd.Player.y
	for i := 0; i < 11; i++ {
		for j := 0; j < 11; j++ {
			if gd.Map[i][j] {
				if checkIfPointInRect(x, y, int32(i)*75, int32(j)*75, 75, 75) {
					return false
				}
			}
		}
	}
	return true
}

func checkIfPossibleRight(gd GameData) bool {
	var x = gd.Player.x + 37
	var y = gd.Player.y
	for i := 0; i < 11; i++ {
		for j := 0; j < 11; j++ {
			if gd.Map[i][j] {
				if checkIfPointInRect(x, y, int32(i)*75, int32(j)*75, 75, 75) {
					return false
				}
			}
		}
	}
	return true
}

// long

func checkIfPossibleUpLong(gd GameData) bool {
	var x = gd.Player.x
	var y = gd.Player.y - 65
	for i := 0; i < 11; i++ {
		for j := 0; j < 11; j++ {
			if gd.Map[i][j] {
				if checkIfPointInRect(x, y, int32(i)*75, int32(j)*75, 75, 75) || checkIfPointInRect(x+30, y, int32(i)*75, int32(j)*75, 75, 75) || checkIfPointInRect(x-30, y, int32(i)*75, int32(j)*75, 75, 75) {
					return false
				}
			}
		}
	}
	return true
}

func checkIfPossibleDownLong(gd GameData) bool {
	var x = gd.Player.x
	var y = gd.Player.y + 65
	for i := 0; i < 11; i++ {
		for j := 0; j < 11; j++ {
			if gd.Map[i][j] {
				if checkIfPointInRect(x, y, int32(i)*75, int32(j)*75, 75, 75) || checkIfPointInRect(x+30, y, int32(i)*75, int32(j)*75, 75, 75) || checkIfPointInRect(x-30, y, int32(i)*75, int32(j)*75, 75, 75) {
					return false
				}
			}
		}
	}
	return true
}

func checkIfPossibleLeftLong(gd GameData) bool {
	var x = gd.Player.x - 65
	var y = gd.Player.y
	for i := 0; i < 11; i++ {
		for j := 0; j < 11; j++ {
			if gd.Map[i][j] {
				if checkIfPointInRect(x, y, int32(i)*75, int32(j)*75, 75, 75) || checkIfPointInRect(x, y+30, int32(i)*75, int32(j)*75, 75, 75) || checkIfPointInRect(x, y-30, int32(i)*75, int32(j)*75, 75, 75) {
					return false
				}
			}
		}
	}
	return true
}

func checkIfPossibleRightLong(gd GameData) bool {
	var x = gd.Player.x + 65
	var y = gd.Player.y
	for i := 0; i < 11; i++ {
		for j := 0; j < 11; j++ {
			if gd.Map[i][j] {
				if checkIfPointInRect(x, y, int32(i)*75, int32(j)*75, 75, 75) || checkIfPointInRect(x, y+30, int32(i)*75, int32(j)*75, 75, 75) || checkIfPointInRect(x, y-30, int32(i)*75, int32(j)*75, 75, 75) {
					return false
				}
			}
		}
	}
	return true
}
