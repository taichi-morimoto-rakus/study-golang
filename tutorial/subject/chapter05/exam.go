package chapter05

func DarumaDrop(daruma []int) []int {
	length := len(daruma)
	if length == 0 || length == 1 {
		return daruma
	}

	mid := length / 2
	if length%2 == 0 {
		if daruma[mid-1] >= daruma[mid] {
			return append(daruma[:mid], daruma[mid+1:]...)
		} else {
			return append(daruma[:mid-1], daruma[mid:]...)
		}
	} else {
		return append(daruma[:mid], daruma[mid+1:]...)
	}
}

func MatrixMultiple(seed []int) [][]int {
	matrix := make([][]int, len(seed))
	for i := range matrix {
		matrix[i] = make([]int, len(seed))
	}

	for i, y := range seed {
		for j, x := range seed {
			matrix[i][j] = y * x
		}
	}
	return matrix
}
