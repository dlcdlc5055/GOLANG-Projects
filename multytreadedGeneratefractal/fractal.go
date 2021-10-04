package main

import (
	"math/cmplx"
	"sync"
)

const MAX_ITERATIONS int = 1750

func calculatePixel(xF, yF float64, wg *sync.WaitGroup, image *[800][800]uint8, x, y int) {
	var z complex128 = 0
	var c complex128 = complex(xF, yF)
	var iterations int = 0

	for iterations < MAX_ITERATIONS {
		z = z*z + c

		if cmplx.Abs(z) > 2 {
			break
		}

		iterations++
	}

	var color uint8 = uint8(256*iterations/MAX_ITERATIONS) * 50

	image[x][y] = color

	wg.Done()
}

func generateFractal(image *[800][800]uint8, done *bool) {
	var wg sync.WaitGroup
	for i := 0; i < 800; i++ {
		for j := 0; j < 800; j++ {
			var xFract float64 = (float64(j) - 800/1.32) * 2.0 / 800
			var yFract float64 = (float64(i) - 800/2) * 2.0 / 800

			wg.Add(1)
			go calculatePixel(xFract, yFract, &wg, image, j, i)
		}
	}
	wg.Wait()
	*done = true
}
