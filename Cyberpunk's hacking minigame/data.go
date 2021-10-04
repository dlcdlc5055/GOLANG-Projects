package main

const horizontal = false
const verical = true

type GameData struct {
	bord         [5][5]string
	bordSubData  [5][5]bool
	correct      [3]string
	buffer       [3]string
	bufferIndex  int
	selectedVal  int
	selectedType bool
	MaxError     int
	Error        int
	corrects     int
}
