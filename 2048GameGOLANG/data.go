package main

import rl "github.com/gen2brain/raylib-go/raylib"

type GameData struct {
	bord        [4][4]int
	score       int
	largestTile int
	win         bool
}

var Color_2048 = rl.Green
var Color_1024 = rl.Purple
var Color_512 = rl.Blue
var Color_256 = rl.Red
var Color_128 = rl.Yellow
var Color_64 = rl.Pink
var Color_32 = rl.DarkPurple
var Color_16 = rl.DarkBrown
var Color_8 = rl.DarkGreen
var Color_4 = rl.Brown
var Color_2 = rl.DarkBlue
