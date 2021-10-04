package main

import "github.com/gen2brain/raylib-go/raylib"
import "strconv"

func Draw(gd gameData) {
	DrawBg()
  DrawBord(gd.bord)
	var str1="White"
	if gd.turn==true{
		str1="Black"
	}
	var nrOfWhitePice=GetNrOfPicesWhite(gd)
	var nrOfBlackPice=GetNrOfPicesBlack(gd)
	if gd.finished==false{
		rl.DrawText("Turn: "+str1,11,805 ,25, rl.White)
		rl.DrawText("White: "+strconv.Itoa( nrOfWhitePice),345,805 ,25, rl.White)
		rl.DrawText("Black: "+strconv.Itoa(nrOfBlackPice),675,805 ,25, rl.White)
	}else{
		rl.DrawText("Press Space To Reset",11,805 ,25, rl.White)
	}
}

func DrawBord(gb GameBord){
  for i := 0; i < 8; i++ {
    for j := 0; j < 8; j++ {
      var color=rl.DarkGreen
      if i%2==1&&j%2==0{
      color=  rl.LightGray
      }
      if i%2==0&&j%2==1{
      color=  rl.LightGray
      }
      if gb.data[i][j]==BlackPice{
        rl.DrawCircle(int32(i)*100+50,int32(j)*100+50, 40,rl.Black)
      }
      if gb.data[i][j]==WhitePice{
        rl.DrawCircle(int32(i)*100+50,int32(j)*100+50, 40,rl.RayWhite)
      }
      if gb.data[i][j]==PosibleMovePice{
        rl.DrawCircle(int32(i)*100+50,int32(j)*100+50, 40,rl.DarkGray)
        rl.DrawCircle(int32(i)*100+50,int32(j)*100+50, 35,color)
      }
      }
    }
}


func DrawBg() {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
      var color=rl.DarkGreen
      if i%2==1&&j%2==0{
      color=  rl.LightGray
      }
      if i%2==0&&j%2==1{
      color=  rl.LightGray
      }
      rl.DrawRectangle(int32(100*i),int32(100*j),100, 100,color)
		}
	}
}
