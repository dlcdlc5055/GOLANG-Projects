package main

type gameData struct {
	bord     GameBord
	state    int
	turn     bool
	selected Vect
	finished bool
	wonSide  int8
}

type GameBord struct {
	data [8][8]int32
}

type Vect struct {
	x int
	y int
}

const PosibleMovePice = 3
const WhitePice = 2
const BlackPice = 1
const EmptyPice = 0
const WhiteTurn = false
const BlackTurn = true
