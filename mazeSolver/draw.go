package main

//import "strconv"
import "github.com/gen2brain/raylib-go/raylib"

func draw(data Data){
  drawMaze(data)
  drawGrid()
  if data.hasSolution==false{
    rl.DrawRectangle(0, 0,900, 900, rl.Black)
    rl.DrawText("Maze Has No Solution", 15, 15, 30.,rl.RayWhite)
  }
}

func drawMaze(data Data){
    for i:=0;i<bordSize;i++{
      for j:=0;j<bordSize;j++{
        if data.buffer[i][j]==10{
          rl.DrawRectangle(int32(i*(900/bordSize)),int32(j*(900/bordSize)),(900/bordSize),(900/bordSize),rl.Blue)
        }
        if data.maze[i][j]==Wall{
          rl.DrawRectangle(int32(i*(900/bordSize)),int32(j*(900/bordSize)),(900/bordSize),(900/bordSize),rl.Black)
        }
        if data.maze[i][j]==start{
          rl.DrawRectangle(int32(i*(900/bordSize)),int32(j*(900/bordSize)),(900/bordSize),(900/bordSize),rl.Green)
        }
        if data.maze[i][j]==end{
            rl.DrawRectangle(int32(i*(900/bordSize)),int32(j*(900/bordSize)),(900/bordSize),(900/bordSize),rl.Red)
        }
        //rl.DrawText(strconv.Itoa(data.mazeSolveBuffer[i][j]), int32(i*(900/bordSize)),int32(j*(900/bordSize)), 20,rl.Black)
      }
    }
}

func drawGrid(){
  for i:=0;i<bordSize;i++{
    rl.DrawRectangle(0, int32(i*(900/bordSize)),900, 1,rl.Black)
    rl.DrawRectangle(int32(i*(900/bordSize)), 0,1, 900,rl.Black)
  }
}
