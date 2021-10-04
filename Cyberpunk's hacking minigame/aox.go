package main

import rl "github.com/gen2brain/raylib-go/raylib"

func generateGD() GameData {
	var gd GameData
	gd.corrects = 0
	//set correct array
	gd.correct[0] = "BD"
	gd.correct[1] = "55"
	gd.correct[2] = "E9"
	//set buffer index
	gd.bufferIndex = 0
	//setbuffer
	for i := 0; i < 3; i++ {
		gd.buffer[i] = ""
	}
	//setbord row1
	gd.bord[0][0] = "BD"
	gd.bord[1][0] = "BD"
	gd.bord[2][0] = "BD"
	gd.bord[3][0] = "E9"
	gd.bord[4][0] = "55"
	//setbord row2
	gd.bord[0][1] = "E9"
	gd.bord[1][1] = "55"
	gd.bord[2][1] = "E9"
	gd.bord[3][1] = "1C"
	gd.bord[4][1] = "55"
	//setbord row3
	gd.bord[0][2] = "55"
	gd.bord[1][2] = "E9"
	gd.bord[2][2] = "BD"
	gd.bord[3][2] = "1C"
	gd.bord[4][2] = "E9"
	//setbord row4
	gd.bord[0][3] = "E9"
	gd.bord[1][3] = "1C"
	gd.bord[2][3] = "1C"
	gd.bord[3][3] = "BD"
	gd.bord[4][3] = "1C"
	//setbord row5
	gd.bord[0][4] = "55"
	gd.bord[1][4] = "E9"
	gd.bord[2][4] = "E9"
	gd.bord[3][4] = "1C"
	gd.bord[4][4] = "55"
	//bord subdata
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			gd.bordSubData[i][j] = true
		}
	}
	//selected
	gd.selectedVal = 0
	gd.selectedType = false
	//error
	gd.Error = 0
	gd.MaxError = 3
	return gd
}

func getMouseX() int {
	var buffer int
	var pos = rl.GetMousePosition()
	if pos.X < 50 || pos.X > 50+85*5 {
		buffer = -1
	} else {
		var bx = pos.X - 50
		buffer = int(bx / 85)
	}
	if buffer > 4 {
		return 4
	}
	return buffer
}

func getMouseY() int {
	var buffer int
	var pos = rl.GetMousePosition()
	if pos.Y < 50 || pos.Y > 50+85*5 {
		buffer = -1
	} else {
		var by = pos.Y - 50
		buffer = int(by / 85)
	}
	if buffer > 4 {
		return 4
	}
	return buffer
}

func checkForWin(gd GameData) bool {
	var count = 0
	for i := 0; i < 3; i++ {
		if gd.correct[i] == gd.buffer[i] {
			count += 1
		}
	}
	if count > 2 {
		return true
	}
	return false
}

func checkForLose(gd GameData) bool {
	if gd.Error > gd.MaxError {
		return true
	}
	return false
}

func checkIfInSelected(gd GameData) bool {
	var x = getMouseX()
	var y = getMouseY()
	if x < 0 || y < 0 {
		return false
	}
	if gd.selectedType == true {
		if x == gd.selectedVal {
			return true
		}
		return false
	} else {
		if y == gd.selectedVal {
			return true
		}
		return false
	}
}
