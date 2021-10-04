package main

type Vect struct {
	x int
	y int
}

type GameData struct {
	active      bool
	poused      bool
	win         bool
	food        Vect
	Snake       []Vect
	SnakeLenght int16
	direction   int8
	bordSize    int16
	score       int16
	timeBuffer  float32
	lose        bool
}

const left = 0
const right = 1
const up = 3
const down = 4
