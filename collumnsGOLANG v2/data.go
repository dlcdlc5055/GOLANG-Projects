package main

type GameData struct {
	bord        [6][13]int8
	pice        Pice
	nextPice    Pice
	score       int
	timeBuffer  float32
	bordPosible [6][13]bool
	mode        bool
}

type Vect struct {
	x int32
	y int32
}

type Pice struct {
	pos     Vect
	content [3]int8
}
