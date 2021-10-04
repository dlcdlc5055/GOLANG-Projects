package main

import (
	"math/rand"
	"time"
)

func generateGameData() GameData {
	var gd GameData
	for i := 0; i < 6; i++ {
		for j := 0; j < 13; j++ {
			gd.bord[i][j] = -1
			gd.bordPosible[i][j] = true
		}
	}
	gd.nextPice = generatePice(0, &gd)
	gd.pice = generatePice(25, &gd)
	return gd
}

func checkIfOcupied(gd GameData, pos Vect, offX, offY int32) bool {
	for i := 0; i < 3; i++ {
		if pos.y+offY+int32(i) > -1 && pos.x+offX < 6 && pos.x+offX > -1 {
			if gd.bord[pos.x+offX][pos.y+offY+int32(i)] > -1 {
				return true
			}
		}
	}
	return false
}

func checkIfPiceInBord(x, y int, gd GameData) bool {
	if x > -1 && x < 6 && y < 13 && y > -1 {
		if gd.bord[x][y] > -1 {
			return true
		} else {
			return false
		}
	}
	return false
}

func swapPice(gd *GameData, x1, y1, x2, y2 int) {
	var buffer = gd.bord[x1][y1]
	gd.bord[x1][y1] = gd.bord[x2][y2]
	gd.bord[x2][y2] = buffer
}

func processFall(gd *GameData) {
	for i := 0; i < 6; i++ {
		processFallRow(gd, i)
	}
}

func processFallRow(gd *GameData, col int) {
	for i := 0; i < 100; i++ {
		for j := 1; j < 13; j++ {
			if gd.bord[col][12-j] > -1 && gd.bord[col][12-j+1] < 0 {
				swapPice(gd, col, 12-j, col, 12-j+1)
			}
		}
	}
}

func ProcessMaches(gd *GameData) {
	var buffer = gd.bord
	//horizontal
	for i := 0; i < 13; i++ {
		ProcessMachesRow(gd, i, buffer)
	}
	//vertical
	for i := 0; i < 6; i++ {
		ProcessMachesCOll(gd, i, buffer)
	}
	//d1
	ProcessMachesd1(gd, buffer)
	//d2
	ProcessMachesd2(gd, buffer)
}

func checkIfBordFull(gd GameData) bool {
	var buffer = 0
	for i := 0; i < 6; i++ {
		for j := 0; j < 13; j++ {
			if gd.bord[i][j] > -1 {
				buffer += 1
			}
		}
	}
	if buffer == 13*6 {
		return true
	}
	return false
}

func ProcessMachesd1(gd *GameData, buffer1 [6][13]int8) {
	var buffer = buffer1
	for i := 0; i < 4; i++ {
		for j := 0; j < 11; j++ {
			if buffer[i][j] == buffer[i+1][j+1] && buffer[i+1][j+1] == buffer[i+2][j+2] && buffer[i][j] > -1 {
				if gd.bord[i][j] > -1 || buffer[i][j] > -1 {
					gd.bord[i][j] = -1
					gd.score += 5
				}
				if gd.bord[i+1][j+1] > -1 || buffer[i+1][j+1] > -1 {
					gd.bord[i+1][j+1] = -1
					gd.score += 5
				}
				if gd.bord[i+2][j+2] > -1 || buffer[i+2][j+2] > -1 {
					gd.bord[i+2][j+2] = -1
					gd.score += 5
				}
			}
		}
	}
}

func ProcessMachesd2(gd *GameData, buffer1 [6][13]int8) {
	var buffer = buffer1
	for i := 0; i < 4; i++ {
		for j := 2; j < 13; j++ {
			if buffer[i][j] == buffer[i+1][j-1] && buffer[i+1][j-1] == buffer[i+2][j-2] && buffer[i][j] > -1 {
				if gd.bord[i][j] > -1 || buffer[i][j] > -1 {
					gd.bord[i][j] = -1
					gd.score += 5
				}
				if gd.bord[i+1][j-1] > -1 || buffer[i+1][j-1] > -1 {
					gd.bord[i+1][j-1] = -1
					gd.score += 5
				}
				if gd.bord[i+2][j-2] > -1 || buffer[i+2][j-2] > -1 {
					gd.bord[i+2][j-2] = -1
					gd.score += 5
				}
			}
		}
	}
}

func ProcessMachesCOll(gd *GameData, coll int, buffer1 [6][13]int8) {
	var buffer = buffer1
	for i := 0; i < 11; i++ {
		if buffer[coll][i] == buffer[coll][i+1] && buffer[coll][i+1] == buffer[coll][i+2] && buffer[coll][i] > -1 {
			if gd.bord[coll][i] > -1 || buffer[coll][i] > -1 {
				gd.bord[coll][i] = -1
				gd.score += 2
			}
			if gd.bord[coll][i+1] > -1 || buffer[coll][i+1] > -1 {
				gd.bord[coll][i+1] = -1
				gd.score += 2
			}
			if gd.bord[coll][i+2] > -1 || buffer[coll][i+2] > -1 {
				gd.bord[coll][i+2] = -1
				gd.score += 2
			}
		}
	}
}

func generatePossiblesForEasyMode(gd *GameData) {
	for i := 0; i < 6; i++ {
		for j := 0; j < 13; j++ {
			gd.bordPosible[i][j] = false
		}
	}
	for i := 0; i < 6; i++ {
		for j := 0; j < 13; j++ {
			generatePossiblesForEasyModePerPice(i, j, gd)
		}
	}
}

func generatePossiblesForEasyModePerPice(x, y int, gd *GameData) {
	generatePossiblesForEasyModePerPiceHorizontal(x, y, gd)
	generatePossiblesForEasyModePerPiceVertical(x, y, gd)
	generatePossiblesForEasyModePerPiceD1(x, y, gd)
	generatePossiblesForEasyModePerPiceD2(x, y, gd)
}

func generatePossiblesForEasyModePerPiceHorizontal(x, y int, gd *GameData) {

}

func generatePossiblesForEasyModePerPiceVertical(x, y int, gd *GameData) {

}

func generatePossiblesForEasyModePerPiceD1(x, y int, gd *GameData) {

}

func generatePossiblesForEasyModePerPiceD2(x, y int, gd *GameData) {

}

func ProcessMachesRow(gd *GameData, row int, buffer1 [6][13]int8) {
	var buffer = buffer1
	for i := 0; i < 4; i++ {
		if buffer[i][row] == buffer[i+1][row] && buffer[i+1][row] == buffer[i+2][row] && buffer[i][row] > -1 {
			if gd.bord[i][row] > -1 || buffer[i][row] > -1 {
				gd.bord[i][row] = -1
				gd.score += 1
			}
			if gd.bord[i+1][row] > -1 || buffer[i+1][row] > -1 {
				gd.bord[i+1][row] = -1
				gd.score += 1
			}
			if gd.bord[i+2][row] > -1 || buffer[i+2][row] > -1 {
				gd.bord[i+2][row] = -1
				gd.score += 1
			}
		}
	}
}

func checkIfPosibleNewPice(gd *GameData) bool {
	var buffer = 0
	for i := 0; i < 6; i++ {
		if gd.bord[i][0] < 0 {
			buffer += 1
		}
	}
	if buffer == 0 {
		return false
	} else {
		return true
	}
}

func PlacePice(gd *GameData) bool {
	var pos = gd.pice.pos
	var content = gd.pice.content
	for i := 0; i < 3; i++ {
		if -1 < pos.y+int32(i) {
			gd.bord[pos.x][pos.y+int32(i)] = content[i]
		}
	}
	var possible = checkIfPosibleNewPice(gd)
	if possible {
		gd.pice = gd.nextPice
		rand.Seed(time.Now().UTC().UnixNano())
		gd.nextPice = generatePice(int64(rand.Intn(100)), gd)
	} else {
		return false
	}
	return true
}

func turnPice(pice Pice) Pice {
	var p2 Pice = pice
	p2.content[0] = pice.content[1]
	p2.content[1] = pice.content[2]
	p2.content[2] = pice.content[0]
	return p2
}

func reverseTurnPice(pice Pice) Pice {
	var p2 Pice = pice
	p2.content[1] = pice.content[0]
	p2.content[2] = pice.content[1]
	p2.content[0] = pice.content[2]
	return p2
}

func generatePice(offset int64, gd *GameData) Pice {
	rand.Seed(time.Now().UTC().UnixNano() + offset)
	var c1 = int8(rand.Intn(typesOfPices))
	var c2 = int8(rand.Intn(typesOfPices))
	var c3 = int8(rand.Intn(typesOfPices))
	var pice Pice
	pice.content[0] = c1
	pice.content[1] = c2
	pice.content[2] = c3
	pice.pos.x = 3
	pice.pos.y = -3
	var value = true
	if gd.bord[3][0] > -1 {
		for value {
			rand.Seed(time.Now().UTC().UnixNano() + offset)
			pice.pos.x = int32(rand.Intn(6))
			if gd.bord[pice.pos.x][0] < 0 {
				value = false
			}
		}
	}
	return pice
}
