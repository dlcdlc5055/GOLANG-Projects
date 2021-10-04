package main

import "math"

import "github.com/gen2brain/raylib-go/raylib"

func CalculateCosArray(value *float64,array *[900]int32,increment *float64){
  for i := 0; i < 900; i++ {
    var cos = math.Cos(*value)
    //println(cos)
    *value += *increment
    //println(((int64(cos*10) + 10) * 12))
    array[i] = ((int32(cos*10) + 10) * 12)
  }

}

func main() {
	var value float64 = 0
	var increment float64 = 0.065
  var buffer=increment
  var incrementModefire = 0.001
	var array [900]int32

  rl.InitWindow(1600, 935, "Cos Function Vizualizer")

  CalculateCosArray(&value, &array, &increment)
rl.SetTargetFPS(60)

for !rl.WindowShouldClose() {
  rl.BeginDrawing()

  if rl.IsKeyDown(rl.KeyW){
    value = 0
    increment+=incrementModefire
  CalculateCosArray(&value, &array, &increment)
  }

  if rl.IsKeyDown(rl.KeyS){
    value = 0
    increment-=incrementModefire
  CalculateCosArray(&value, &array, &increment)
  }

  if rl.IsKeyPressed(rl.KeySpace){
    value = 0
    increment=buffer
  CalculateCosArray(&value, &array, &increment)
  }

  rl.ClearBackground(rl.RayWhite)

  DrawCosFunctionArray(array)


rl.DrawFPS(10,10)
rl.DrawRectangle(0, 33, 1600, 4,rl.DarkBlue)

  rl.EndDrawing()
}

rl.CloseWindow()
}

func DrawCosFunctionArray(array [900]int32){
  	for i := 0; i < 900; i++ {
      rl.DrawRectangle(0, int32(i)+35,1600,1, rl.Color{0,uint8(array[i]),0,255})
    }
}
