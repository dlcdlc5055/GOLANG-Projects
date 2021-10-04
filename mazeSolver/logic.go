package main

func generateData() Data {
	var data = Data{}
	data.hasEnd = false
	data.hasStart = false
	data.state = 0
	data.hasSolution = true
	for i := 0; i < bordSize; i++ {
		for j := 0; j < bordSize; j++ {
			data.maze[i][j] = 0
			data.mazeSolveBuffer[i][j] = 0
		}
	}
	return data
}

func solveMaze(data *Data) {
	var buffer = false
	if data.hasStart == true && data.hasEnd == true {
		for i := 0; i < bordSize; i++ {
			for j := 0; j < bordSize; j++ {
				if data.maze[i][j] == Wall {
					data.buffer[i][j] = Wall
				} else {
					data.buffer[i][j] = 0
				}
			}
		}
		floodfill(1, data, data.StartX, data.StartY, &buffer)
		for i := 0; i < bordSize; i++ {
			for j := 0; j < bordSize; j++ {
				if data.maze[i][j] == Wall {
					data.buffer[i][j] = Wall
				} else {
					data.buffer[i][j] = 0
				}
			}
		}
		data.hasSolution = buffer
		var b1 int = 0
		generatePath(data.EndX, data.EndY, data, &b1)
	}
}

func generatePath(x, y int, data *Data, b1 *int) {
	data.buffer[x][y] = 10
	if data.mazeSolveBuffer[x][y] == 0 {
		return
	}
	*b1 = data.mazeSolveBuffer[x][y]
	//x
	if checkifinbounds(x-1, y) && data.mazeSolveBuffer[x-1][y] < *b1 {
		generatePath(x-1, y, data, b1)
	}
	if checkifinbounds(x+1, y) && data.mazeSolveBuffer[x+1][y] < *b1 {
		generatePath(x+1, y, data, b1)
	}
	//y
	if checkifinbounds(x, y-1) && data.mazeSolveBuffer[x][y-1] < *b1 {
		generatePath(x, y-1, data, b1)
	}
	if checkifinbounds(x, y+1) && data.mazeSolveBuffer[x][y+1] < *b1 {
		generatePath(x, y+1, data, b1)
	}
}

func checkifinbounds(x, y int) bool {
	if x < 0 || x > 16 || y < 0 || y > 16 {
		return false
	}
	return true
}

func floodfill(value int, data *Data, x, y int, buffer *bool) {
	println(value)
	var old = 0
	var new = 10
	if x < 0 || x >= 17 || y < 0 || y >= 17 {
		return
	}
	if data.buffer[x][y] != old {
		return
	}

	if data.mazeSolveBuffer[x][y] != 0 {
		return
	}

	data.buffer[x][y] = new
	data.mazeSolveBuffer[x][y] = value
	if x == data.EndX && y == data.EndY {
		*buffer = true
	}

	floodfill(value+1, data, x+1, y, buffer)
	floodfill(value+1, data, x, y+1, buffer)
	floodfill(value+1, data, x-1, y, buffer)
	floodfill(value+1, data, x, y-1, buffer)
}
