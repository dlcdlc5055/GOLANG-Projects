package main

import "github.com/gen2brain/raylib-go/raylib"

func Update(gd *gameData) {
	var buffer int32 = WhitePice
	if gd.turn == BlackTurn {
		buffer = BlackPice
	}
  var isPressed=rl.IsMouseButtonPressed(rl.MouseLeftButton)
  var x = GetSqered(int(rl.GetMouseX()))
  var y = GetSqered(int(rl.GetMouseY()))
  var b2=CheckIfPosibleCordInBord(x, y)
  if b2==true && isPressed{
    if gd.bord.data[x][y]==PosibleMovePice{
      gd.bord.data[x][y]=buffer
      ProcessChanges(x,y,gd)
      gd.turn=!gd.turn
    }
  }
	ClearPosibleMove(gd)
	GeneratePosibleMovesForSide(gd)
	CheckIfFinished(gd)
	if gd.finished{
		if rl.IsKeyPressed(rl.KeySpace){
			*gd=initGameData()
		}
	}
}
