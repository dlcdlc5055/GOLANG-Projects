package main

const Block int = 1
const Win int = 2

const left = 1
const right = 2
const up = 3
const down = 4

type GameData struct {
	Map              [11][10]int
	Enemys           [2]Vect
	EnemysVect       [2]int
	Player           Vect
	Items            [3]Vect
	PlayerTimeBuffer float32
	EnemyTimeBuffer  float32
	Poused           bool
	ItemsActive      [3]bool
	ColectedItems    int
	PlayerVect       int
}

type Vect struct {
	x int32
	y int32
}
