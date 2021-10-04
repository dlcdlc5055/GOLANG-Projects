package main

func initGameData() gameData {
	var gd = gameData{}
	gd.state = 1
	gd.turn = WhiteTurn
	gd.bord = initGameBord()
	gd.selected = Vect{-1, -1}
	gd.finished = false
	return gd
}

func initGameBord() GameBord {
	var gb = GameBord{}
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			gb.data[i][j] = EmptyPice
		}
	}
	gb.data[3][3] = BlackPice
	gb.data[4][4] = BlackPice
	gb.data[4][3] = WhitePice
	gb.data[3][4] = WhitePice
	return gb
}

func GeneratePosibleMovesForSide(gd *gameData) {
	var buffer int32 = WhitePice
	if gd.turn == BlackTurn {
		buffer = BlackPice
	}
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if buffer == gd.bord.data[i][j] {
				GeneratePosibleMovesForPice(i, j, gd)
			}
		}
	}
}

func ClearPosibleMove(gd *gameData) {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if gd.bord.data[i][j] == PosibleMovePice {
				gd.bord.data[i][j] = EmptyPice
			}
		}
	}
}

func GeneratePosibleMovesForPice(x, y int, gd *gameData) {
	var buffer = gd.bord.data[x][y]
	var enemy = WhitePice
	if buffer == WhitePice {
		enemy = BlackPice
	}
	if CheckIfPosibleCordInBord(x, y+1) && gd.bord.data[x][y+1] == int32(enemy) && CheckIfPosibleCordInBord(x, y+2) && gd.bord.data[x][y+2] == 0 {
		gd.bord.data[x][y+2] = PosibleMovePice
	}
	if CheckIfPosibleCordInBord(x, y-1) && gd.bord.data[x][y-1] == int32(enemy) && CheckIfPosibleCordInBord(x, y-2) && gd.bord.data[x][y-2] == 0 {
		gd.bord.data[x][y-2] = PosibleMovePice
	}
	if CheckIfPosibleCordInBord(x+1, y) && gd.bord.data[x+1][y] == int32(enemy) && CheckIfPosibleCordInBord(x+2, y) && gd.bord.data[x+2][y] == 0 {
		gd.bord.data[x+2][y] = PosibleMovePice
	}
	if CheckIfPosibleCordInBord(x-1, y) && gd.bord.data[x-1][y] == int32(enemy) && CheckIfPosibleCordInBord(x-2, y) && gd.bord.data[x-2][y] == 0 {
		gd.bord.data[x-2][y] = PosibleMovePice
	}

	if CheckIfPosibleCordInBord(x+1, y+1) && gd.bord.data[x+1][y+1] == int32(enemy) && CheckIfPosibleCordInBord(x+2, y+2) && gd.bord.data[x+2][y+2] == 0 {
		gd.bord.data[x+2][y+2] = PosibleMovePice
	}
	if CheckIfPosibleCordInBord(x-1, y-1) && gd.bord.data[x-1][y-1] == int32(enemy) && CheckIfPosibleCordInBord(x-2, y-2) && gd.bord.data[x-2][y-2] == 0 {
		gd.bord.data[x-2][y-2] = PosibleMovePice
	}
	if CheckIfPosibleCordInBord(x-1, y+1) && gd.bord.data[x-1][y+1] == int32(enemy) && CheckIfPosibleCordInBord(x-2, y+2) && gd.bord.data[x-2][y+2] == 0 {
		gd.bord.data[x-2][y+2] = PosibleMovePice
	}
	if CheckIfPosibleCordInBord(x+1, y-1) && gd.bord.data[x+1][y-1] == int32(enemy) && CheckIfPosibleCordInBord(x+2, y-2) && gd.bord.data[x+2][y-2] == 0 {
		gd.bord.data[x+2][y-2] = PosibleMovePice
	}
}

func CheckIfPosibleCordInBord(x, y int) bool {
	return x < 8 && y < 8 && x > -1 && y > -1
}

func ProcessChanges(x, y int, gd *gameData) {
	println("Executed")
	var buffer int = int(gd.bord.data[x][y])
	var enemy int = WhitePice
	if buffer == WhitePice {
		enemy = BlackPice
	}
	var b1, b2, b3, b4, v1, v2, v3, v4 bool
	var left, right, up, down, vert1, vert2, vert3, vert4 int
	b1 = true
	b2 = true
	b3 = true
	b4 = true
	v1 = true
	v2 = true
	v3 = true
	v4 = true
	vert1 = 1
	vert2 = 1
	vert3 = 1
	vert4 = 1
	left = 1
	right = 1
	up = 1
	down = 1
	for i := 1; i < 8; i++ {
		//left
		if CheckIfPosibleCordInBord(x-i, y) && gd.bord.data[x-i][y] == int32(enemy) && b1 {
			left += 1
		} else {
			b1 = false
		}
		//right
		if CheckIfPosibleCordInBord(x+i, y) && gd.bord.data[x+i][y] == int32(enemy) && b2 {
			right += 1
		} else {
			b2 = false
		}
		//up
		if CheckIfPosibleCordInBord(x, y-i) && gd.bord.data[x][y-i] == int32(enemy) && b3 {
			up += 1
		} else {
			b3 = false
		}
		//down
		if CheckIfPosibleCordInBord(x, y+i) && gd.bord.data[x][y+i] == int32(enemy) && b4 {
			down += 1
		} else {
			b4 = false
		}
	}

	//verct
	for i := 1; i < 8; i++ {
		if CheckIfPosibleCordInBord(x-i, y-i) && gd.bord.data[x-i][y-i] == int32(enemy) && v1 {
			vert1 += 1
		} else {
			v1 = false
		}

		if CheckIfPosibleCordInBord(x+i, y+i) && gd.bord.data[x+i][y+i] == int32(enemy) && v2 {
			vert2 += 1
		} else {
			v2 = false
		}

		if CheckIfPosibleCordInBord(x-i, y+i) && gd.bord.data[x-i][y+i] == int32(enemy) && v3 {
			vert3 += 1
		} else {
			v3 = false
		}

		if CheckIfPosibleCordInBord(x+i, y-i) && gd.bord.data[x+i][y-i] == int32(enemy) && v4 {
			vert4 += 1
		} else {
			v4 = false
		}
	}

	println(left)
	println(right)
	println(up)
	println(down)
	println(vert1)
	println(vert2)
	println(vert3)
	println(vert4)

	//left
	if x-left > -1 && gd.bord.data[x-left][y] == int32(buffer) {
		println("left")
		for i := 0; i < left; i++ {
			gd.bord.data[x-i][y] = int32(buffer)
		}
	}
	//right
	if x+right < 8 && gd.bord.data[x+right][y] == int32(buffer) {
		println("right")
		for i := 0; i < right; i++ {
			gd.bord.data[x+i][y] = int32(buffer)
		}
	}
	//up
	if y-up > -1 && gd.bord.data[x][y-up] == int32(buffer) {
		println("up")
		for i := 0; i < up; i++ {
			gd.bord.data[x][y-i] = int32(buffer)
		}
	}
	//down
	if y+down < 8 && gd.bord.data[x][y+down] == int32(buffer) {
		println("down")
		for i := 0; i < down; i++ {
			gd.bord.data[x][y+i] = int32(buffer)
		}
	}
	//vvvv
	if x-vert1 > -1 && y-vert1 > -1 && gd.bord.data[x-vert1][y-vert1] == int32(buffer) {
		println("Vect1")
		for i := 0; i < vert1; i++ {
			gd.bord.data[x-i][y-i] = int32(buffer)
		}
	}
	if x+vert2 < 8 && y+vert2 < 8 && gd.bord.data[x+vert2][y+vert2] == int32(buffer) {
		println("Vect2")
		for i := 0; i < vert2; i++ {
			gd.bord.data[x+i][y+i] = int32(buffer)
		}
	}
	if x-vert3 > -1 && y+vert3 < 8 && gd.bord.data[x-vert3][y+vert3] == int32(buffer) {
		println("Vect3")
		for i := 0; i < vert3; i++ {
			gd.bord.data[x-i][y+i] = int32(buffer)
		}
	}
	if x+vert4 < 8 && y-vert4 > -1 && gd.bord.data[x+vert4][y-vert4] == int32(buffer) {
		println("Vect4")
		for i := 0; i < vert4; i++ {
			gd.bord.data[x+i][y-i] = int32(buffer)
		}
	}
}

func GetSqered(value int) int {
	return value / 100
}

func GetNrOfPicesWhite(gd gameData) int {
	var count int = 0
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if gd.bord.data[i][j] == WhitePice {
				count++
			}
		}
	}
	return count
}

func GetNrOfPosibleMoves(gd gameData) int {
	var count int = 0
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if gd.bord.data[i][j] == PosibleMovePice {
				count++
			}
		}
	}
	return count
}

func GetNrOfPicesBlack(gd gameData) int {
	var count int = 0
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if gd.bord.data[i][j] == BlackPice {
				count++
			}
		}
	}
	return count
}

func CheckIfBordFull(gd gameData) bool {
	var white = GetNrOfPicesWhite(gd)
	var black = GetNrOfPicesBlack(gd)
	if (white + black) == 64 {
		return true
	} else {
		return false
	}
}

func CheckIfFinished(gd *gameData) {
	if CheckIfBordFull(*gd) {
		gd.finished = true
	}
}
