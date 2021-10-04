package main

func generateGD() GameData {
	var gd GameData

	gd.PlayerVect = 0

	for i := 0; i < 11; i++ {
		gd.Map[i][0] = Block
		gd.Map[i][1] = Block
		gd.Map[i][9] = Block
	}
	gd.Map[1][1] = Win
	for i := 0; i < 10; i++ {
		gd.Map[0][i] = Block
		gd.Map[10][i] = Block
	}
	gd.Map[2][3] = Block
	gd.Map[2][4] = Block
	gd.Map[2][5] = Block
	gd.Map[4][5] = Block
	gd.Map[8][5] = Block
	gd.Map[2][6] = Block
	gd.Map[2][7] = Block

	gd.Map[4][3] = Block
	gd.Map[5][3] = Block
	gd.Map[6][3] = Block
	gd.Map[7][3] = Block
	gd.Map[8][3] = Block

	gd.Map[6][5] = Block
	gd.Map[6][6] = Block
	gd.Map[6][7] = Block
	gd.Map[5][7] = Block
	gd.Map[4][7] = Block
	gd.Map[7][7] = Block
	gd.Map[8][7] = Block

	gd.Poused = false

	gd.Enemys[0].x = 1
	gd.Enemys[0].y = 2
	gd.Enemys[1].x = 9
	gd.Enemys[1].y = 8

	gd.Items[0].x = 1
	gd.Items[0].y = 5
	gd.Items[1].x = 5
	gd.Items[1].y = 4
	gd.Items[2].x = 9
	gd.Items[2].y = 5

	gd.ItemsActive[0] = true
	gd.ItemsActive[1] = true
	gd.ItemsActive[2] = true

	gd.ColectedItems = 0

	gd.Player.x = 7
	gd.Player.y = 4
	return gd
}

func checkIfPossibleEnemy(val1, val2 int32, gd GameData, i int) bool {
	if gd.Map[val1+gd.Enemys[i].x][val2+gd.Enemys[i].y] != 0 {
		return false
	}
	return true
}

func checkIfPossiblePlayer(val1, val2 int32, gd GameData) bool {
	if gd.Map[val1+gd.Player.x][val2+gd.Player.y] == Block {
		return false
	}
	return true
}

func getifItemPick(gd GameData) int {
	for i := 0; i < 3; i++ {
		if gd.ItemsActive[i] {
			var x = gd.Items[i].x == gd.Player.x
			var y = gd.Items[i].y == gd.Player.y
			if x && y {
				return i + 1
			}
		}
	}
	return 0
}

func checkForWin(gd GameData) bool {
	var x = gd.Player.x
	var y = gd.Player.y
	if gd.Map[x][y] == Win {
		return true
	}
	return false
}

func checkForLose(gd GameData) bool {
	for i := 0; i < 2; i++ {
		var x = gd.Player.x == gd.Enemys[i].x
		var y = gd.Player.y == gd.Enemys[i].y
		if x && y {
			return true
		}
	}
	return false
}
