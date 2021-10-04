package main

import "math/rand"

func MoveToLeft(gd *GameData) bool {
	var bufferGd = *gd
	//move
	for v := 0; v < 25; v++ {
		for j := 0; j < 4; j++ {
			if gd.bord[0][j] == 0 && gd.bord[1][j] != 0 {
				swap(0, j, 1, j, gd)
			} else if gd.bord[1][j] == 0 && gd.bord[2][j] != 0 {
				swap(1, j, 2, j, gd)
			} else if gd.bord[2][j] == 0 && gd.bord[3][j] != 0 {
				swap(2, j, 3, j, gd)
			}
		}
	}
	//combine
	for j := 0; j < 4; j++ {
		if gd.bord[0][j] == gd.bord[1][j] {
			gd.bord[0][j] *= 2
			gd.bord[1][j] = 0
		}
		if gd.bord[1][j] == gd.bord[2][j] {
			gd.bord[1][j] *= 2
			gd.bord[2][j] = 0
		}
		if gd.bord[2][j] == gd.bord[3][j] {
			gd.bord[2][j] *= 2
			gd.bord[3][j] = 0
		}
	}
	//move
	for v := 0; v < 25; v++ {
		for j := 0; j < 4; j++ {
			if gd.bord[0][j] == 0 && gd.bord[1][j] != 0 {
				swap(0, j, 1, j, gd)
			} else if gd.bord[1][j] == 0 && gd.bord[2][j] != 0 {
				swap(1, j, 2, j, gd)
			} else if gd.bord[2][j] == 0 && gd.bord[3][j] != 0 {
				swap(2, j, 3, j, gd)
			}
		}
	}
	println(CheckIfIdenticleBord(bufferGd, *gd))
	if !CheckIfIdenticleBord(bufferGd, *gd) {
		return true
	} else {
		return false
	}
}

func MoveToRight(gd *GameData) bool {
	var bufferGd = *gd
	//move
	for v := 0; v < 25; v++ {
		for j := 0; j < 4; j++ {
			if gd.bord[1][j] == 0 && gd.bord[0][j] != 0 {
				swap(1, j, 0, j, gd)
			}
			if gd.bord[2][j] == 0 && gd.bord[1][j] != 0 {
				swap(2, j, 1, j, gd)
			}
			if gd.bord[3][j] == 0 && gd.bord[2][j] != 0 {
				swap(3, j, 2, j, gd)
			}
		}
	}
	//combine
	for j := 0; j < 4; j++ {
		if gd.bord[3][j] == gd.bord[2][j] {
			gd.bord[3][j] *= 2
			gd.bord[2][j] = 0
		} else if gd.bord[2][j] == gd.bord[1][j] {
			gd.bord[2][j] *= 2
			gd.bord[1][j] = 0
		} else if gd.bord[1][j] == gd.bord[0][j] {
			gd.bord[1][j] *= 2
			gd.bord[0][j] = 0
		}
	}
	//move
	for v := 0; v < 25; v++ {
		for j := 0; j < 4; j++ {
			if gd.bord[1][j] == 0 && gd.bord[0][j] != 0 {
				swap(1, j, 0, j, gd)
			}
			if gd.bord[2][j] == 0 && gd.bord[1][j] != 0 {
				swap(2, j, 1, j, gd)
			}
			if gd.bord[3][j] == 0 && gd.bord[2][j] != 0 {
				swap(3, j, 2, j, gd)
			}
		}
	}
	println(CheckIfIdenticleBord(bufferGd, *gd))
	if !CheckIfIdenticleBord(bufferGd, *gd) {
		return true
	} else {
		return false
	}
}

func MoveToUp(gd *GameData) bool {
	var bufferGd = *gd
	//move
	for v := 0; v < 25; v++ {
		for j := 0; j < 4; j++ {
			if gd.bord[j][0] == 0 && gd.bord[j][1] != 0 {
				swap(j, 0, j, 1, gd)
			}
			if gd.bord[j][1] == 0 && gd.bord[j][2] != 0 {
				swap(j, 1, j, 2, gd)
			}
			if gd.bord[j][2] == 0 && gd.bord[j][3] != 0 {
				swap(j, 2, j, 3, gd)
			}
		}
	}
	//combine
	for j := 0; j < 4; j++ {
		if gd.bord[j][0] == gd.bord[j][1] {
			gd.bord[j][0] *= 2
			gd.bord[j][1] = 0
		}
		if gd.bord[j][1] == gd.bord[j][2] {
			gd.bord[j][1] *= 2
			gd.bord[j][2] = 0
		}
		if gd.bord[j][2] == gd.bord[j][3] {
			gd.bord[j][2] *= 2
			gd.bord[j][3] = 0
		}
	}
	//move
	for v := 0; v < 25; v++ {
		for j := 0; j < 4; j++ {
			if gd.bord[j][0] == 0 && gd.bord[j][1] != 0 {
				swap(j, 0, j, 1, gd)
			} else if gd.bord[j][1] == 0 && gd.bord[j][2] != 0 {
				swap(j, 1, j, 2, gd)
			} else if gd.bord[j][2] == 0 && gd.bord[j][3] != 0 {
				swap(j, 2, j, 3, gd)
			}
		}
	}
	println(CheckIfIdenticleBord(bufferGd, *gd))
	if !CheckIfIdenticleBord(bufferGd, *gd) {
		return true
	} else {
		return false
	}
}

func MoveToDown(gd *GameData) bool {
	var bufferGd = *gd
	//move
	for v := 0; v < 25; v++ {
		for j := 0; j < 4; j++ {
			if gd.bord[j][1] == 0 && gd.bord[j][0] != 0 {
				swap(j, 1, j, 0, gd)
			} else if gd.bord[j][2] == 0 && gd.bord[j][1] != 0 {
				swap(j, 2, j, 1, gd)
			} else if gd.bord[j][3] == 0 && gd.bord[j][2] != 0 {
				swap(j, 3, j, 2, gd)
			}
		}
	}
	//combine
	for j := 0; j < 4; j++ {
		if gd.bord[j][3] == gd.bord[j][2] {
			gd.bord[j][3] *= 2
			gd.bord[j][2] = 0
		}
		if gd.bord[j][2] == gd.bord[j][1] {
			gd.bord[j][2] *= 2
			gd.bord[j][1] = 0
		}
		if gd.bord[j][1] == gd.bord[j][0] {
			gd.bord[j][1] *= 2
			gd.bord[j][0] = 0
		}
	}
	//move
	for v := 0; v < 25; v++ {
		for j := 0; j < 4; j++ {
			if gd.bord[j][1] == 0 && gd.bord[j][0] != 0 {
				swap(j, 1, j, 0, gd)
			} else if gd.bord[j][2] == 0 && gd.bord[j][1] != 0 {
				swap(j, 2, j, 1, gd)
			} else if gd.bord[j][3] == 0 && gd.bord[j][2] != 0 {
				swap(j, 3, j, 2, gd)
			}
		}
	}
	println(CheckIfIdenticleBord(bufferGd, *gd))
	if !CheckIfIdenticleBord(bufferGd, *gd) {
		return true
	} else {
		return false
	}
}

func GenerateNewTile(gd *GameData) {
	var FullTiles = getNrOfTiles(gd)
	if FullTiles < 16 {
		for true {
			var x = rand.Intn(4)
			var y = rand.Intn(4)
			if gd.bord[x][y] == 0 {
				var is4 = rand.Intn(5)
				if is4 > 3 {
					gd.bord[x][y] = 4
				} else {
					gd.bord[x][y] = 2
				}
				break
			}
		}
	}
}
