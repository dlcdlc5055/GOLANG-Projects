package main

import "github.com/gen2brain/raylib-go/raylib"
import "strconv"

func render(gd GameData) {
	if !gd.lose {
		DrawFood(gd.food, int32(gd.bordSize))
		DrawSnake(gd.Snake, int32(gd.bordSize), gd.SnakeLenght)
		DrawScore(gd.score)
	} else {
		rl.DrawText("Score: "+strconv.Itoa(int(gd.score)), 10, 10, 30, rl.RayWhite)
		rl.DrawText("Press Space To Reset", 90, 435, 63, rl.RayWhite)
	}
}

func DrawSnake(snake []Vect, size int32, snakeLenght int16) {
	for i := len(snake) - int(snakeLenght); i < len(snake); i++ {
		rl.DrawRectangle(int32(snake[i].x)*size, int32(snake[i].y)*size, size, size, rl.RayWhite)
	}
}

func DrawFood(food Vect, size int32) {
	rl.DrawRectangle(int32(food.x)*size, int32(food.y)*size, size, size, rl.Red)
}

func DrawScore(score int16) {
	rl.DrawRectangle(0, 900, 10000, 10000, rl.RayWhite)
	rl.DrawText("Score: "+strconv.Itoa(int(score)), 10, 905, 40, rl.Black)
}
