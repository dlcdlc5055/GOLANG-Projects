package main

func generateGD() GameData {
	var gd GameData
	gd.PoliceMoved = false
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			var block Block
			block.left = false
			block.right = false
			block.up = false
			block.down = false
			gd.Map[i][j] = block
		}
	}
	for i := 0; i < 8; i++ {
		if i != 5 {
			gd.Map[i][0].down = true
		}
	}
	gd.Map[5][0].left = true
	gd.Map[5][0].right = true
	for i := 1; i < 9; i++ {
		gd.Map[8][i].left = true
	}
	gd.Finish.X = 5
	gd.Finish.Y = 0
	gd.Player.X = 1
	gd.Player.Y = 2
	gd.Police.X = 1
	gd.Police.Y = 1
	gd.Map[1][2].up = true
	gd.Map[2][2].up = true
	gd.Map[2][3].left = true
	gd.Map[2][4].left = true
	gd.Map[2][4].down = true
	gd.Map[1][4].down = true
	gd.Map[0][6].down = true
	gd.Map[3][5].right = true
	gd.Map[3][6].right = true
	gd.Map[3][7].right = true
	gd.Map[3][7].right = true
	gd.Map[5][3].left = true
	gd.Map[5][3].right = true
	gd.Map[5][3].down = true
	gd.Map[5][7].down = true
	gd.Map[7][3].up = true
	gd.Poused = false
	for i := 0; i < 4; i++ {
		gd.Map[7][4+i].left = true
	}
	return gd
}

func checkForLose(gd GameData) bool {
	if gd.Player.X == gd.Police.X && gd.Player.Y == gd.Police.Y {
		return true
	}
	return false
}

func checkForWin(gd GameData) bool {
	if gd.Player.X == gd.Finish.X && gd.Player.Y == gd.Finish.Y {
		return true
	}
	return false
}

func reset(gd *GameData) {
	*gd = generateGD()
}
