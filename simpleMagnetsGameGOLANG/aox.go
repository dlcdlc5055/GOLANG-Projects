package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func checkIfWin(gd GameData) bool {
	var buffer = 0
	for i := 0; i < 8; i++ {
		if gd.correct[i] == gd.player[i] {
			buffer += 1
		}
	}
	if buffer == 8 {
		return true
	}
	return false
}

func getDiference(v1, v2 int32) int32 {
	if v1 < v2 {
		return v2 - v1
	}
	if v1 > v2 {
		return v1 - v2
	}
	return 0
}

func calculateDistance(x1, y1, x2, y2 int32) int32 {
	return int32(math.Sqrt(float64(getDiference(x1, x2)*getDiference(x1, x2) + getDiference(y1, y2)*getDiference(y1, y2))))
}

func getClickedBtn() int {
	var data int = -1
	for i := 0; i < 7; i++ {
		if calculateDistance(200, 50+int32(i)*90+80, int32(rl.GetMousePosition().X), int32(rl.GetMousePosition().Y)) < 30 && rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			data = i
		}
	}
	return data
}
