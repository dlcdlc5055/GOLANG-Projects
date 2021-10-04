package main

import rl "github.com/gen2brain/raylib-go/raylib"

func update(gd *GameData) {
	if rl.IsKeyPressed(rl.KeySpace) {
		reset(gd)
	}
	var win = checkForWin(*gd)
	var lose = checkForLose(*gd)
	if win {
		drawWin()
		gd.Poused = true
	} else if lose {
		drawLose()
		gd.Poused = true
	} else {
	}
	if !gd.Poused {
		ProcessPlayer(gd)
	}
}

func ProcessPlayer(gd *GameData) {
	var block = gd.Map[gd.Player.X][gd.Player.Y]
	var x = gd.Player.X
	var y = gd.Player.Y
	if rl.IsKeyPressed(rl.KeyW) {
		if y == 0 {
			return
		}
		if block.up == false && gd.Map[gd.Player.X][y-1].down == false {
			if checkForLose(*gd) == false || checkForWin(*gd) == false {
				y -= 1
				if y < 0 {
					y = 0
				}
				if y > 8 {
					y = 8
				}
				gd.Player.X = x
				gd.Player.Y = y
				gd.PoliceMoved = false
				ProcessAI(gd)
			}

		}
	}
	if rl.IsKeyPressed(rl.KeyS) {
		if y == 8 {
			return
		}
		if block.down == false && gd.Map[gd.Player.X][y+1].up == false {
			if checkForLose(*gd) == false || checkForWin(*gd) == false {
				y += 1
				if y < 0 {
					y = 0
				}
				if y > 8 {
					y = 8
				}
				gd.Player.X = x
				gd.Player.Y = y
				gd.PoliceMoved = false
				ProcessAI(gd)

			}

		}
	}
	if rl.IsKeyPressed(rl.KeyA) {
		if x == 0 {
			return
		}
		if block.left == false && gd.Map[gd.Player.X-1][y].right == false {
			if checkForLose(*gd) == false || checkForWin(*gd) == false {
				x -= 1
				if x < 0 {
					x = 0
				}
				if x > 8 {
					x = 8
				}
				gd.Player.X = x
				gd.Player.Y = y
				gd.PoliceMoved = false
				ProcessAI(gd)

			}
		}
	}
	if rl.IsKeyPressed(rl.KeyD) {
		if x == 8 {
			return
		}
		if block.right == false && gd.Map[gd.Player.X+1][y].left == false {
			if checkForLose(*gd) == false || checkForWin(*gd) == false {
				x += 1
				if x < 0 {
					x = 0
				}
				if x > 8 {
					x = 8
				}
				gd.Player.X = x
				gd.Player.Y = y
				gd.PoliceMoved = false
				ProcessAI(gd)

			}

		}
	}
	if x < 0 {
		x = 0
	}
	if x > 8 {
		x = 8
	}
	if y < 0 {
		y = 0
	}
	if y > 8 {
		y = 8
	}
	//pollice error fix
	if gd.Police.X < 0 {
		gd.Police.X = 0
	}
	if gd.Police.Y < 0 {
		gd.Police.Y = 0
	}
	if gd.Police.X > 8 {
		gd.Police.X = 8
	}
	if gd.Police.Y > 8 {
		gd.Police.Y = 8
	}
	gd.Player.X = x
	gd.Player.Y = y
}

func GenerateBlockForAI(gd *GameData) Block {
	var block Block
	var x = gd.Police.X
	var y = gd.Police.Y
	block.left = gd.Map[x][y].left
	block.right = gd.Map[x][y].right
	block.up = gd.Map[x][y].up
	block.down = gd.Map[x][y].down
	if x > 0 {
		var blockNew = gd.Map[x-1][y]
		if blockNew.right {
			block.left = true
		}
	}
	if x < 8 {
		var blockNew = gd.Map[x+1][y]
		if blockNew.left {
			block.right = true
		}
	}
	if y > 0 {
		var blockNew = gd.Map[x][y-1]
		if blockNew.down {
			block.up = true
		}
	}
	if y < 8 {
		var blockNew = gd.Map[x][y+1]
		if blockNew.up {
			block.down = true
		}
	}
	return block
}

const Horizontal = false
const Vertical = true

func ProcessAI(gd *GameData) {
	gd.PoliceMoved = false
	gd.PoliceBUffer.X = gd.Police.X
	gd.PoliceBUffer.Y = gd.Police.Y
	if gd.Player.X == gd.Police.X && gd.Player.Y == gd.Police.Y {
		return
	}
	if gd.Player.X != gd.Police.X && gd.Player.Y != gd.Police.Y {
		var block = GenerateBlockForAI(gd)
		var b1 = ProcessAIMoves(gd, block, false)
		block = GenerateBlockForAI(gd)
		var b2 = ProcessAIMoves(gd, block, true)
		if b2 == false || b1 == false {
			block = GenerateBlockForAI(gd)
			ProcessAIMoves(gd, block, false)
		} else if b2 == false && b1 == false {
			var block = GenerateBlockForAI(gd)
			ProcessAIMoves(gd, block, true)
			block = GenerateBlockForAI(gd)
			var val2 = ProcessAIMoves(gd, block, false)
			if val2 == true && !gd.PoliceMoved {
				gd.PoliceMoved = true
				gd.PoliceBUffer.X = gd.Police.X
				gd.PoliceBUffer.Y = gd.Police.Y
			}
		}
	} else if gd.Player.Y == gd.Police.Y {
		var block = GenerateBlockForAI(gd)
		ProcessAIMoves(gd, block, false)
		block = GenerateBlockForAI(gd)
		ProcessAIMoves(gd, block, false)
	} else if gd.Player.X == gd.Police.X {
		var block = GenerateBlockForAI(gd)
		ProcessAIMoves(gd, block, true)
		block = GenerateBlockForAI(gd)
		ProcessAIMoves(gd, block, true)
	}
	gd.PoliceMoved = true
}

func ProcessAIMoves(gd *GameData, block Block, axis bool) bool {
	if checkForWin(*gd) {
		return true
	}
	if checkForLose(*gd) {
		return true
	}
	if axis == Horizontal {
		if gd.Player.X < gd.Police.X && !block.left {
			gd.Police.X -= 1
			return true
		} else if gd.Player.X > gd.Police.X && !block.right {
			gd.Police.X += 1
			return true
		}
	} else if axis == Vertical {
		if gd.Player.Y < gd.Police.Y && !block.up {
			gd.Police.Y -= 1
			return true
		} else if gd.Player.Y > gd.Police.Y && !block.down {
			gd.Police.Y += 1
			return true
		}
	}
	return false
}
