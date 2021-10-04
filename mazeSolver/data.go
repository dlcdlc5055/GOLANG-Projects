package main

const Wall = 1
const end = 2
const start = 3
const path = 4

type Data struct {
	StartX          int
	StartY          int
	EndX            int
	EndY            int
	hasEnd          bool
	hasStart        bool
	state           int
	buffer          [bordSize][bordSize]int
	maze            [bordSize][bordSize]int
	mazeSolveBuffer [bordSize][bordSize]int
	hasSolution     bool
}
