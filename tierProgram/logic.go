package main

import (
	"math"
	"math/rand"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func getClickedCircleIndex(circles []Circle, xMouse, yMouse int32) int {
	var length = len(circles)
	for i := 0; i < length; i++ {
		if CirclePointCollision(circles[i].x, circles[i].y, xMouse, yMouse, circles[i].radius+4, false) {
			return i
		}
	}
	return -1
}

func CirclePointCollision(xCircle, yCircle, xMouse, yMouse int32, radius float64, tp bool) bool {
	var cx, cy int32
	cx = diference(xCircle, xMouse)
	cy = diference(yCircle, yMouse)
	var bufferX, bufferY int32
	bufferX = cx * cx
	bufferY = cy * cy
	var buffer int32 = int32(math.Sqrt(float64(bufferX + bufferY)))
	if float64(buffer) < radius && !tp {
		return true
	} else if float64(buffer) < (radius*2) && tp {
		return true
	}
	return false
}

func checkCircleCollisionForCircleArray(xCircle, yCircle int32, radius float64, circles []Circle) bool {
	var length = len(circles)
	for i := 0; i < length; i++ {
		if CirclePointCollision(xCircle, yCircle, circles[i].x, circles[i].y, circles[i].radius+8, true) {
			return true
		}
	}
	return false
}

func GenerateCircles(nrOfCircles int) []Circle {
	rand.Seed(time.Now().UTC().UnixNano())
	var circles []Circle
	var buffer = 0
	var colors = []rl.Color{rl.Red, rl.Orange, rl.Green, rl.Yellow, rl.Blue, rl.Purple, rl.Pink}
	var count = 0
	for buffer < nrOfCircles {
		count += 1
		var x = int32(rand.Intn(890 - 110))
		var y = int32(rand.Intn(900 - 110))
		var colorIndex = rand.Int() % len(colors)
		var circle = Circle{x + 965, y + 55, 50, colors[colorIndex], true, false}
		if !checkCircleCollisionForCircleArray(circle.x, circle.y, 50, circles) && getNumberOfCircleWithColor(colors[colorIndex], circles) < 7 {
			circles = append(circles, circle)
			buffer += 1
		}
		if count > 150000 {
			break
		}
	}
	return circles
}

func getNrOfActiveCircles(circles []Circle) int {
	var count = 0
	for i := 0; i < len(circles); i++ {
		if circles[0].active {
			count += 1
		}
	}
	return count
}

func getNumberOfCircleWithColor(color rl.Color, circles []Circle) int {
	var count = 0
	for i := 0; i < len(circles); i++ {
		if circles[i].color == color {
			count += 1
		}
	}
	return count
}

func getOffSetForCircle(m1, m2 float32, c1, c2 int32, vectOffset *Vector) {
	var d1 = diference(int32(m1), c1)
	var d2 = diference(int32(m2), c2)
	if c1 < int32(m1) {
		d1 = d1 * -1
	}
	if c2 < int32(m2) {
		d2 = d2 * -1
	}
	vectOffset.x = d1
	vectOffset.y = d2
}

func diference(val1, val2 int32) int32 {
	var buffer int32
	if val1 < val2 {
		buffer = val2 - val1
	} else if val1 > val2 {
		buffer = val1 - val2
	} else {
		buffer = 0
	}
	return buffer
}

type Circle struct {
	x      int32
	y      int32
	radius float64
	color  rl.Color
	active bool
	done   bool
}

type Vector struct {
	x int32
	y int32
}

func backupCircles(c1 []Circle, b1 *[]Circle) {
	*b1 = nil
	for val := 0; val < len(c1); val++ {
		*b1 = append(*b1, (c1)[val])
	}
}

func reset(circles *[]Circle, selectedCircleIndex *int, isSelectedCircle *bool, vectOffset *Vector, DoneCirclesCount *[7]int8, DoneCirclesColors *[7][7]rl.Color, circlesBackUp *[]Circle) {
	*circles = GenerateCircles(20)
	backupCircles(*circles, circlesBackUp)
	*selectedCircleIndex = -1
	*isSelectedCircle = false
	for i := 0; i < 7; i++ {
		(*DoneCirclesCount)[i] = 0
		for j := 0; j < 7; j++ {
			(*DoneCirclesColors)[i][j] = rl.White
		}
	}
}

func CheckForWin(circles []Circle) bool {
	var colors = []rl.Color{rl.Red, rl.Orange, rl.Yellow, rl.Green, rl.Blue, rl.Purple, rl.Pink}
	var CorrectCount = 0
	for i := 0; i < len(circles); i++ {
		if circles[i].x < 900 && circles[i].x > 100 {
			var row = circles[i].y / 128
			if colors[row] == circles[i].color {
				CorrectCount += 1
			}
		}
	}
	if CorrectCount == len(circles) {
		return true
	} else {
		return false
	}
}

func RemoveCircle(x, y int, DoneCirclesCount *[7]int8, DoneCirclesColors *[7][7]rl.Color, circles *[]Circle) {
	(*DoneCirclesCount)[y] -= 1
	var b1 = int32(138 + 50 + 10 + x*108)
	var b2 = int32(y*128 + 50 + 12)
	for i := x + 1; i < 7; i++ {
		DoneCirclesColors[i-1][y] = DoneCirclesColors[i][y]
	}
	for i := 0; i < len(*circles); i++ {
		if (*circles)[i].y == b2 && (*circles)[i].done == true && b1 < (*circles)[i].x {
			var b1 = int32(138 + 50 + 10 + x*108)
			var b2 = int32(y*128 + 50 + 12)
			if (*circles)[i].x == b1 && (*circles)[i].y == b2 {
				(*circles)[i].active = false
			}
			(*circles)[i].x -= 108
		}
	}
}
