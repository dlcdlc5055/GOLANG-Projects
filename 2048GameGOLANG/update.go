package main

func update(gd *GameData) {
	ProcessInput(gd)
	gd.largestTile = GetLargestTile(*gd)
	gd.score = GetScore(*gd)
	gd.win = CheckIfWin(*gd)
}
