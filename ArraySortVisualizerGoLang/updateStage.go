package main

func updateStage(b1, b2, active *int, array *[400]int) {
	if !CheckIfSorted(array) {
		for i := 0; i < 399; i++ {
			if array[i] > array[i+1] {
				*b1 = i
				*b2 = i + 1
				swap(*b1, *b2, array)
				break
			}
		}
	}
}

func MultiStageUpdate(b1, b2, active *int, array *[400]int, stages int) {
	for i := 0; i < stages; i++ {
		updateStage(b1, b2, active, array)
	}
}
