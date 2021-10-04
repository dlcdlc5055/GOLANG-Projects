package main

import "github.com/gen2brain/raylib-go/raylib"
func update(data *Data) {
  var i=getCord(int(rl.GetMouseX()))
  var j=getCord(int(rl.GetMouseY()))
  if rl.IsKeyPressed(rl.KeyEnter){
    *data=generateData()
  }
  if(data.state==0 && i!=-1 && j!=-1){
    if  rl.IsKeyPressed(rl.KeySpace){
      data.state=1
      solveMaze(data)
    }
    if rl.IsKeyPressed(rl.KeyEnter){
      *data=generateData()
    }
    //generate walls
    if(rl.IsMouseButtonDown(rl.MouseLeftButton)&&data.maze[i][j]!=Wall){
      data.maze[i][j]=Wall
    }
    if(rl.IsMouseButtonDown(rl.MouseLeftButton)&&data.maze[i][j]==Wall && rl.IsKeyDown(rl.KeyLeftShift)){
      data.maze[i][j]=0
    }
    //generate start and end
    if(rl.IsMouseButtonDown(rl.MouseRightButton)&&!rl.IsKeyPressed(rl.KeyLeftShift) && data.maze[i][j]==0){
      if data.hasStart==false{
        data.hasStart=true
        data.StartX=i
        data.StartY=j
        data.maze[i][j]=start
      }else if data.hasEnd==false{
        data.hasEnd=true
        data.EndX=i
        data.EndY=j
        data.maze[i][j]=end
      }
    }else if(rl.IsMouseButtonDown(rl.MouseRightButton)&& rl.IsKeyPressed(rl.KeyLeftShift)){
      if data.maze[i][j]==start{
        data.hasStart=false
        data.maze[i][j]=0
      }else if data.maze[i][j]==end{
        data.hasEnd=false
        data.maze[i][j]=0
      }
    }
  }
}

func getCord(value int)int{
  if value<(-17) || value>900-17{
    return -1
  }else{
    return value/(900/bordSize)
  }
}
