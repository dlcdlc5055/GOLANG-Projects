package main

import "github.com/gen2brain/raylib-go/raylib"

func main() {
rl.InitWindow(800, 830, "Analog Clock GOLANG")

rl.SetTargetFPS(60)

var gd=initGameData()

GeneratePosibleMovesForSide(&gd)

for !rl.WindowShouldClose() {

 rl.BeginDrawing()

 Update(&gd)

 rl.ClearBackground(rl.Black)

 Draw(gd)

 rl.EndDrawing()
}

rl.CloseWindow()
}
