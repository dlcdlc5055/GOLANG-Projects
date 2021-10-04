package main

type GameData struct {
	Player       Vect
	Police       Vect
	PoliceBUffer Vect
	PoliceMoved  bool
	Finish       Vect
	Poused       bool
	Map          [9][9]Block
}

type Block struct {
	left  bool
	right bool
	up    bool
	down  bool
}

type Vect struct {
	X int32
	Y int32
}
