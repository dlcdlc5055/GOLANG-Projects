package main

type GameData struct {
	Map             [11][11]bool
	Items           [3]Vect
	ItemsActive     [3]bool
	ItemsCollected  int
	Player          Vect
	Enemys          [2]Vect
	PlayerMoveVect  int
	EnemysMoveVects [2]int
	Lose            bool
	Win             bool
	poused          bool
	AiTimeBuffer    float32
	Active          bool
}

type Vect struct {
	x int32
	y int32
}

const left = 1
const right = 2
const up = 3
const down = 4
