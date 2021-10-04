package main

import (
	"math/rand"
	"time"
)

func ArrayGenerate() [400]int {
	var array [400]int
	for i := 0; i < 400; i++ {
		array[i] = (i * 2) + 50
	}
	ShuffleArray(&array)
	return array
}

func ShuffleArray(array *[400]int) [400]int {
	rand.Seed(time.Now().UTC().UnixNano())
	var count = 0
	for true {
		count++
		var x1 = rand.Intn(400)
		var x2 = rand.Intn(400)
		swap(x1, x2, array)
		if count >= 10000 {
			break
		}
	}
	return *array
}
