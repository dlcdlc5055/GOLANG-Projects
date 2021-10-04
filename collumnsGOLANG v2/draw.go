package main

import (
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const blockSize int32 = 62

func draw(gd GameData, paused bool) {
	var colors []rl.Color
	colors = append(colors, rl.Red)
	colors = append(colors, rl.Blue)
	colors = append(colors, rl.Green)
	colors = append(colors, rl.Yellow)
	colors = append(colors, rl.Orange)
	colors = append(colors, rl.Purple)
	drawBG()
	drawNextPice(gd, colors)
	drawScore(gd)
	drawBord(gd, colors)
	drawActivePice(gd, colors)
	drawMode(gd)
	drawPaused(paused)
}

func drawPaused(paused bool) {
	var text = ""
	if paused {
		text = "paused"
	} else {
		text = "active"
	}
	rl.DrawRectangle(850, blockSize*13-50, 150, 50, rl.Black)
	rl.DrawText(text, 855, blockSize*13-45, 40, rl.RayWhite)
}

func drawActivePice(gd GameData, colors []rl.Color) {
	var pice = gd.pice
	var pos = pice.pos
	for i := 0; i < 3; i++ {
		var value = pice.content[i]
		if pice.pos.y+int32(i) > -1 {
			if pice.content[i]%2 == 1 {
				rl.DrawCircle(1200/2-3*blockSize+(62/2)+int32(pos.x)*blockSize, 50+(62/2)+int32(pos.y+int32(i))*blockSize, 23, colors[value])
			} else {
				rl.DrawRectangle(1200/2-3*blockSize+10+int32(pos.x)*blockSize, 50+10+int32(pos.y+int32(i))*blockSize, 62-20, 62-20, colors[value])
			}
		}
	}
}

func drawBord(gd GameData, colors []rl.Color) {
	var bord = gd.bord
	var bordPossible = gd.bordPosible
	for i := 0; i < 6; i++ {
		for j := 0; j < 13; j++ {
			var value = bord[i][j]
			var value2 = bordPossible[i][j]
			if !gd.mode {
				value2 = true
			}
			if value2 {
				if value > -1 {
					if value%2 == 1 {
						rl.DrawCircle(1200/2-3*blockSize+(62/2)+int32(i)*blockSize, 50+(62/2)+int32(j)*blockSize, 23, colors[value])
					} else {
						rl.DrawRectangle(1200/2-3*blockSize+10+int32(i)*blockSize, 50+10+int32(j)*blockSize, 62-20, 62-20, colors[value])
					}
				}
			} else {
				if value > -1 {
					if value%2 == 1 {
						rl.DrawCircle(1200/2-3*blockSize+(62/2)+int32(i)*blockSize, 50+(62/2)+int32(j)*blockSize, 23, rl.Gray)
					} else {
						rl.DrawRectangle(1200/2-3*blockSize+10+int32(i)*blockSize, 50+10+int32(j)*blockSize, 62-20, 62-20, rl.Gray)
					}
				}
			}
		}
	}
}

func drawScore(gd GameData) {
	var score = gd.score
	var scoreTXT = strconv.Itoa(score)
	rl.DrawText("Score:"+scoreTXT, 60, 408, 50, rl.RayWhite)
}

func drawNextPice(gd GameData, colors []rl.Color) {
	var data = gd.nextPice.content
	if data[0]%2 == 1 {
		rl.DrawCircle(50+blockSize/2, 50+blockSize/2, 23, colors[data[0]])
	} else {
		rl.DrawRectangle(60, 60, 62-20, 62-20, colors[data[0]])
	}
	if data[1]%2 == 1 {
		rl.DrawCircle(50+blockSize/2, 50+blockSize/2+62, 23, colors[data[1]])
	} else {
		rl.DrawRectangle(60, 60+62, 62-20, 62-20, colors[data[1]])
	}
	if data[2]%2 == 1 {
		rl.DrawCircle(50+blockSize/2, 50+blockSize/2+62+62, 23, colors[data[2]])
	} else {
		rl.DrawRectangle(60, 60+62+62, 62-20, 62-20, colors[data[2]])
	}
}

func drawBG() {
	rl.DrawRectangle(0, 0, 10000, 10000, rl.DarkGreen)
	rl.DrawRectangle(1200/2-3*blockSize, 50, blockSize*6, blockSize*13, rl.Black)

	for i := 0; i < 5; i++ {
		rl.DrawLine(1200/2-3*blockSize+int32(i+1)*blockSize, 50, 1200/2-3*blockSize+int32(i+1)*blockSize, 854, rl.Gray)
	}

	for j := 0; j < 12; j++ {
		rl.DrawLine(1200/2-3*blockSize, 50+int32(j+1)*blockSize, 1200/2+3*blockSize, 50+int32(j+1)*blockSize, rl.Gray)
	}

	rl.DrawRectangle(50, 50, blockSize, blockSize*3, rl.Black)

	rl.DrawLine(50, 50+blockSize, 50+blockSize, 50+blockSize, rl.Gray)
	rl.DrawLine(50, 50+blockSize+blockSize, 50+blockSize, 50+blockSize+blockSize, rl.Gray)

	rl.DrawRectangle(50, 400, 320, 60, rl.Black)

	rl.DrawLine(1200/2-3*blockSize, 50, 1200/2+3*blockSize, 50, rl.Gray)
	rl.DrawLine(1200/2-3*blockSize, 50+blockSize*13, 1200/2+3*blockSize, 50+blockSize*13, rl.Gray)
	rl.DrawLine(1200/2-3*blockSize, 50, 1200/2-3*blockSize, 50+blockSize*13, rl.Gray)
	rl.DrawLine(1200/2+3*blockSize, 50, 1200/2+3*blockSize, 50+blockSize*13, rl.Gray)

	rl.DrawRectangle(850, 50, 220, 50, rl.Black)
}

func drawMode(gd GameData) {
	var text = ""
	if gd.mode {
		text = "Mone:Easy"
	} else {
		text = "Mone:Hard"
	}
	rl.DrawText(text, 855, 55, 40, rl.RayWhite)
}
