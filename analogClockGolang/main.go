package main

import "math"
import "time"
import "github.com/gen2brain/raylib-go/raylib"

func main() {
rl.InitWindow(900, 900, "Analog Clock GOLANG")

rl.SetTargetFPS(60)

var bgImage rl.Image=*rl.LoadImage("assets/bg.png")
var bgTexture=rl.LoadTextureFromImage(&bgImage)
var CurrentHour=time.Now().Hour()%12
var CurrentMinute=time.Now().Minute()
var CurrentSecond=time.Now().Second()

println(CurrentHour)
println(CurrentMinute)
println(CurrentSecond)

for !rl.WindowShouldClose() {

 CurrentHour=time.Now().Hour()%12
 CurrentMinute=time.Now().Minute()
 CurrentSecond=time.Now().Second()
 rl.BeginDrawing()

 rl.ClearBackground(rl.Gray)

 rl.DrawCircle(450,450, 400,rl.RayWhite)

 rl.DrawTexture(bgTexture,50,50,rl.RayWhite)

//hour hand
 drawClockHand(230, angleToRad(float64(CurrentHour*-30-(CurrentMinute/2)-180)),25, rl.Black)

//minute hand
 drawClockHand(285, angleToRad(float64(CurrentMinute*-6)+180),20, rl.Black)

 //second hand
  drawClockHand(330, angleToRad(float64(CurrentSecond*-6)+180),15, rl.Red)

 rl.DrawCircle(450,450,33,rl.Black)

 rl.EndDrawing()
}

rl.CloseWindow()
}

func angleToRad(angle float64) float64{
  return angle * (rl.Pi/180)
}

func drawClockHand(lenght float64,angle_rad float64,thicknes float32,color rl.Color){
  var cos float64=math.Cos(angle_rad)
  var sin float64=math.Sin(angle_rad)
  rl.DrawLineEx(rl.Vector2{450,450}, rl.Vector2{450+float32(sin*lenght),450+float32(cos*lenght)}, thicknes, color)
  rl.DrawCircle(int32(450+float32(sin*lenght)),int32(450+float32(cos*lenght)),thicknes/2,color)
}
