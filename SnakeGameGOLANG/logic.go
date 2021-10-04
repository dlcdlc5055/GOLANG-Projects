package main

import (
	"math/rand"
	"time"
)

func initGD(gd *GameData) {
	rand.Seed(time.Now().UTC().UnixNano())
	gd.poused = false
	gd.active = true
	gd.lose = false
	gd.win = false
	gd.Snake = gd.Snake[:0]
	gd.SnakeLenght = 2
	gd.Snake = append(gd.Snake, Vect{rand.Intn(int(28)) + 1, rand.Intn(int(28)) + 1})
	gd.Snake = append(gd.Snake, Vect{gd.Snake[0].x + 1, gd.Snake[0].y})
	gd.direction = right
	gd.bordSize = 900 / 30
	gd.food = Vect{rand.Intn(int(30)), rand.Intn(int(30))}
	for true {
		if gd.food.x != gd.Snake[0].x && gd.food.y != gd.Snake[0].y && gd.food.x != gd.Snake[1].x && gd.food.y != gd.Snake[1].y {
			break
		} else {
			gd.food = Vect{rand.Intn(int(28)) + 1, rand.Intn(int(28)) + 1}
		}
	}
	gd.score = 0
}

func Diference(val1, val2 int) int {
	if val1 == val2 {
		return 0
	} else if val1 < val2 {
		return val2 - val1
	} else {
		return val1 - val2
	}
}

func CheckIfInSnake(Point Vect, gd GameData) bool {
	for i := len(gd.Snake) - int(gd.SnakeLenght); i < len(gd.Snake); i++ {
		if Diference(gd.Snake[i].x, Point.x) == 0 {
			if Diference(gd.Snake[i].y, Point.y) == 0 {
				return true
			}
		}
	}
	return false
}
