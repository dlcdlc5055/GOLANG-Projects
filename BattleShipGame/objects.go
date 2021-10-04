package main

type GameData struct {
	bordRed               [10][10]int
	bordBlue              [10][10]int
	sunkenShipsPlayerRed  [5]bool
	sunkenShipsPlayerBlue [5]bool
	placedShipsPlayerRed  [5]bool
	placedShipsPlayerBlue [5]bool
	state                 int
	isGameStarted         bool
	destroyerBlue         [2]Vect
	destroyerRed          [2]Vect
	submarineBlue         [3]Vect
	submarineRed          [3]Vect
	cruserBlue            [3]Vect
	cruserRed             [3]Vect
	BattleshipBlue        [4]Vect
	BattleshipRed         [4]Vect
	carrierBlue           [5]Vect
	carrierRed            [5]Vect
}

type Vect struct {
	x int32
	y int32
}

const ship = 1
const hitShip = 2
const miss = 3
