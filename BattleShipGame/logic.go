package main

func InitGameData() GameData {
	var gd = GameData{}
	gd.isGameStarted = false
	for i := 0; i < 5; i++ {
		gd.placedShipsPlayerBlue[i] = false
		gd.placedShipsPlayerRed[i] = false
		gd.sunkenShipsPlayerRed[i] = false
		gd.sunkenShipsPlayerBlue[i] = false
	}
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			gd.bordBlue[i][j] = 0
			gd.bordRed[i][j] = 0
		}
	}
	gd.state = 1
	return gd
}
