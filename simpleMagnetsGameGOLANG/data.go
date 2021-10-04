package main

import (
	"math/rand"
	"time"
)

type GameData struct {
	player  [8]bool
	correct [8]bool
	nrMoves int
}

func generateGameData() GameData {
	var gd GameData
	gd.nrMoves = 0
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < 8; i++ {
		gd.player[i] = false
		gd.correct[i] = false
	}
	for i := 0; i < rand.Intn(4)+5; i++ {
		var buffer = rand.Intn(7)
		gd.correct[buffer] = !gd.correct[buffer]
		gd.correct[buffer+1] = !gd.correct[buffer+1]
	}
	var val = true
	var buffer = 0
	var b2 = 0
	for val {
		rand.Seed(time.Now().UTC().UnixNano())
		buffer = 0
		for i := 0; i < 8; i++ {
			if gd.correct[0] == gd.correct[i] {
				buffer += 1
			}
			if gd.correct[i] == false {
				b2 += 1
			}
		}
		if buffer == 8 {
			for i := 0; i < rand.Intn(7)+5; i++ {
				var buffer = rand.Intn(7)
				gd.correct[buffer] = !gd.correct[buffer]
				gd.correct[buffer+1] = !gd.correct[buffer+1]
			}
		} else if b2 > 3 {
			val = false
		}
	}
	return gd
}
