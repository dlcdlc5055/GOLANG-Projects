package main

import (
	"math/rand"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const PlayerSpeed float32 = .6
const EnemySpeed float32 = .85

func update(gd *GameData) {
	var win = checkForWin(*gd)
	var lose = checkForLose(*gd)

	if win {
		drawWin(*gd)
	} else if lose {
		drawLose()
	}

	if !win && !lose {
		var itemPicked = getifItemPick(*gd)
		if itemPicked != 0 {
			gd.ItemsActive[itemPicked-1] = false
			gd.ColectedItems += 1
		}

		if gd.PlayerVect != 0 {
			ProcessEnemy(gd, EnemySpeed)
			EnemyAI(gd, EnemySpeed)
		}

		if gd.PlayerTimeBuffer < PlayerSpeed {
			if gd.PlayerVect == 0 {
				if rl.IsKeyPressed(rl.KeyW) {
					if checkIfPossiblePlayer(0, -1, *gd) {
						gd.PlayerVect = up
					}
				}
				if rl.IsKeyPressed(rl.KeyS) {
					if checkIfPossiblePlayer(0, 1, *gd) {
						gd.PlayerVect = down
					}
				}
				if rl.IsKeyPressed(rl.KeyA) {
					if checkIfPossiblePlayer(1, 0, *gd) {
						gd.PlayerVect = left
					}
				}
				if rl.IsKeyPressed(rl.KeyD) {
					if checkIfPossiblePlayer(-1, 0, *gd) {
						gd.PlayerVect = right
					}
				}
			} else {
				if rl.IsKeyDown(rl.KeyW) {
					if checkIfPossiblePlayer(0, -1, *gd) && gd.PlayerVect != down {
						gd.PlayerVect = up
					}
				}
				if rl.IsKeyDown(rl.KeyS) {
					if checkIfPossiblePlayer(0, 1, *gd) && gd.PlayerVect != up {
						gd.PlayerVect = down
					}
				}
				if rl.IsKeyDown(rl.KeyA) {
					if checkIfPossiblePlayer(-1, 0, *gd) && gd.PlayerVect != right {
						gd.PlayerVect = left
					}
				}
				if rl.IsKeyDown(rl.KeyD) {
					if checkIfPossiblePlayer(1, 0, *gd) && gd.PlayerVect != left {
						gd.PlayerVect = right
					}
				}
			}
		}

		updateplayer(gd, PlayerSpeed)
	}
}

func ProcessEnemy(gd *GameData, delay float32) {
	gd.EnemyTimeBuffer += rl.GetFrameTime()
	if gd.EnemyTimeBuffer > delay {
		for i := 0; i < 2; i++ {
			var dir = gd.EnemysVect[i]
			if dir == up && checkIfPossibleEnemy(0, -1, *gd, i) {
				gd.Enemys[i].y -= 1
			} else if dir == down && checkIfPossibleEnemy(0, 1, *gd, i) {
				gd.Enemys[i].y += 1
			} else if dir == left && checkIfPossibleEnemy(-1, 0, *gd, i) {
				gd.Enemys[i].x -= 1
			} else if dir == right && checkIfPossibleEnemy(1, 0, *gd, i) {
				gd.Enemys[i].x += 1
			}
		}
	}
}

func EnemyAI(gd *GameData, delay float32) {
	if gd.EnemyTimeBuffer > delay {
		for i := 0; i < 2; i++ {
			for true {
				rand.Seed(time.Now().UnixNano())
				var dir = rand.Intn(4) + 1
				if dir == up && gd.EnemysVect[i] != down && checkIfPossibleEnemy(0, -1, *gd, i) {
					gd.EnemysVect[i] = dir
					break
				} else if dir == down && gd.EnemysVect[i] != up && checkIfPossibleEnemy(0, 1, *gd, i) {
					gd.EnemysVect[i] = dir
					break
				} else if dir == right && gd.EnemysVect[i] != left && checkIfPossibleEnemy(1, 0, *gd, i) {
					gd.EnemysVect[i] = dir
					break
				} else if dir == left && gd.EnemysVect[i] != right && checkIfPossibleEnemy(-1, 0, *gd, i) {
					gd.EnemysVect[i] = dir
					break
				}
			}
		}
		gd.EnemyTimeBuffer = 0
	}
}

func updateplayer(gd *GameData, delay float32) {
	gd.PlayerTimeBuffer += rl.GetFrameTime()
	if gd.PlayerTimeBuffer > delay {
		gd.PlayerTimeBuffer = 0
		if gd.PlayerVect == up && checkIfPossiblePlayer(0, -1, *gd) {
			gd.Player.y -= 1
		}
		if gd.PlayerVect == left && checkIfPossiblePlayer(-1, 0, *gd) {
			gd.Player.x -= 1
		}
		if gd.PlayerVect == right && checkIfPossiblePlayer(1, 0, *gd) {
			gd.Player.x += 1
		}
		if gd.PlayerVect == down && checkIfPossiblePlayer(0, 1, *gd) {
			gd.Player.y += 1
		}
	}
}
