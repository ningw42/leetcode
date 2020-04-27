func setZeroes(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				defer setZero(matrix, i, j)
			}
		}
	}
}

func setZero(matrix [][]int, i, j int) {
	for x := 0; x < len(matrix); x++ {
		matrix[x][j] = 0
	}
	for y := 0; y < len(matrix[0]); y++ {
		matrix[i][y] = 0
	}
}