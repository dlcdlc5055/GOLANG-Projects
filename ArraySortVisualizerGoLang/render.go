package main

import "github.com/gen2brain/raylib-go/raylib"

func renderArray(b1, b2 int, array [400]int) {
	for i := 0; i < 400; i++ {
		if i != b2 && i != b2 {
			rl.DrawRectangle(int32(i*4)+1, int32(900-array[i]), 2, 1000000, rl.RayWhite)
		} else if i == b1 {
			rl.DrawRectangle(int32(i*4)+1, int32(900-array[i]), 2, 1000000, rl.Green)
		} else if i == b2 {
			rl.DrawRectangle(int32(i*4)+1, int32(900-array[i]), 2, 1000000, rl.Red)
		}
	}
}
