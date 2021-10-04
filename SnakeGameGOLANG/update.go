package main

import (
	"math/rand"

	"github.com/gen2brain/raylib-go/raylib"
)

func update(gd *GameData) {
	if !gd.lose {
		ProcessInput(gd)
		updateTime(gd)
		updateSnake(gd)
	} else {
		if rl.IsKeyPressed(rl.KeySpace) {
			initGD(gd)
		}
	}
}

func updateSnake(gd *GameData) {
	if gd.timeBuffer > 0.18 {
		gd.timeBuffer = 0
		var offsetX = 0
		var offsetY = 0
		if gd.direction == right {
			offsetX = 1
			offsetY = 0
		} else if gd.direction == left {
			offsetX = -1
			offsetY = 0
		} else if gd.direction == down {
			offsetX = 0
			offsetY = +1
		} else if gd.direction == up {
			offsetX = 0
			offsetY = -1
		}
		var point = Vect{gd.Snake[len(gd.Snake)-1].x + offsetX, gd.Snake[len(gd.Snake)-1].y + offsetY}
		if point.x > 29 {
			point.x = 0
		}
		if point.y > 29 {
			point.y = 0
		}
		if point.x < 0 {
			point.x = 29
		}
		if point.y < 0 {
			point.y = 29
		}
		if !CheckIfInSnake(point, *gd) {
			gd.Snake = append(gd.Snake, point)
		} else {
			gd.lose = true
		}
		if CheckIfInSnake(gd.food, *gd) {
			gd.score += 1
			gd.SnakeLenght += 1
			for true {
				if !CheckIfInSnake(gd.food, *gd) {
					break
				} else {
					gd.food = Vect{rand.Intn(int(28)) + 1, rand.Intn(int(28)) + 1}
				}
			}
		}
	}
}

func updateTime(gd *GameData) {
	gd.timeBuffer += rl.GetFrameTime()
}
