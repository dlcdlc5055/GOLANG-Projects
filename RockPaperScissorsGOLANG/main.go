package main

import "github.com/gen2brain/raylib-go/raylib"
import "math/rand"
import "time"

const Rock=0
const Paper=1
const Scissors=2

func main() {
	rl.InitWindow(800, 900, "Rock Paper Scissors")

  var assetsImage = rl.LoadImage("assets/assets.png")
  var assetsTexture = rl.LoadTextureFromImage(assetsImage)
  var selectedTop int=3
  var selectedBottom int=-1
  var isSelected bool = false
  var text string = ""
  var textX int32 =280
	var color=rl.Black

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

    DrawAssets(assetsTexture)
    DrawCovers(selectedTop,selectedBottom)

    if text=="You Win!!"{
      textX=285
			color=rl.Green
    }
    if text=="You Lose!!"{
      textX=265
			color=rl.Red
    }
    if text=="Draw!!"{
      textX=320
			color=rl.Black
    }


    rl.DrawText(text,textX,410,50,color)

    CheckForSelect(&selectedBottom, &isSelected,&selectedTop)

    if isSelected{
      if rl.IsKeyPressed(rl.KeySpace){
        reset(&selectedBottom,&isSelected,&selectedTop,&text)
      }
    }

    checkForWin(&text, selectedBottom, isSelected, selectedTop)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}

func checkForWin(text *string,selectedBottom int,isSelected bool,selectedTop int){
  if selectedBottom==Rock{
    if selectedTop==Scissors{
      *text="You Win!!"
    }else if selectedTop==Paper{
      *text="You Lose!!"
    }else if selectedTop==Rock{
      *text="Draw!!"
    }
  }
  if selectedBottom==Scissors{
    if selectedTop==Scissors{
      *text="Draw!!"
    }else if selectedTop==Paper{
      *text="You Win!!"
    }else if selectedTop==Rock{
      *text="You Lose!!"
    }
  }
  if selectedBottom==Paper{
    if selectedTop==Scissors{
      *text="You Lose!!"
    }else if selectedTop==Paper{
      *text="Draw!!"
    }else if selectedTop==Rock{
      *text="You Win!!"
    }
  }
}

func reset(selectedBottom *int,isSelected *bool,selectedTop *int,text *string){
  *text=""
  *selectedBottom=-1
  *selectedTop=-1
  *isSelected=false
}

func CheckForSelect(selectedBottom *int,isSelected *bool,selectedTop *int){
  if *isSelected{
    return
  }
  if rl.IsMouseButtonPressed(rl.MouseLeftButton){
    if rl.GetMouseY()>500{
      var x = rl.GetMouseX()
      if x>0 && x<800{
        *isSelected=true
        *selectedTop=randSelect()
        *selectedBottom=int(x/(800/3))
      }
    }
  }
}

func randSelect() int{
  rand.Seed(time.Now().UnixNano())
  var value=rand.Intn(3)
  return value
}

func DrawCovers(s1,s2 int){
  if s1!=-1{
    for i:=0;i<3; i++{
      if i!=s1{
        rl.DrawRectangle((800/3)*int32(i), 0,(800/3),400,rl.RayWhite)
      }
    }
  }
  if s2!=-1{
    for i:=0;i<3; i++{
      if i!=s2{
        rl.DrawRectangle((800/3)*int32(i), 500,(800/3),400,rl.RayWhite)
      }
    }
  }
}

func DrawAssets(assetsTexture rl.Texture2D){
  rl.DrawTexture(assetsTexture, 0,0, rl.RayWhite)

  rl.DrawTexture(assetsTexture, 0,510, rl.RayWhite)
}
