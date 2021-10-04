package main

type GameData struct {
	bord          [25][18]bool
	actives       [5]bool
	state         int
	selectedVect  [4]Vect
	selectedIndex int
	score         int
}

type Vect struct {
	x int
	y int
}
