func rotate(matrix [][]int) {
	var stash int
	for i := 0; i < len(matrix) / 2; i++ {
		// stash top
		for j := len(matrix)-1-i; j > i; j-- {
			stash = matrix[i][j]
			matrix[i][j] = matrix[len(matrix)-1-j][i]
			matrix[len(matrix)-1-j][i] = matrix[len(matrix)-1-i][len(matrix)-1-j]
			matrix[len(matrix)-1-i][len(matrix)-1-j] = matrix[j][len(matrix)-1-i]
			matrix[j][len(matrix)-1-i] = stash
		}
	}
}

func naive(matrix [][]int) {
	for i := 0; i < len(matrix) / 2; i++ {
		// stash top
		tmp := make([]int, len(matrix) - 2 * i)
		for j := len(matrix)-1-i; j >= i; j-- {
			tmp[j-i] = matrix[i][j]
		}
		// rotate left to top
		for j := len(matrix)-1-i; j >= i; j-- {
			matrix[i][j] = matrix[len(matrix)-1-j][i]
		}
		// rotate bottom to left
		for j := len(matrix)-1-i; j >= i; j-- {
			matrix[len(matrix)-1-j][i] = matrix[len(matrix)-1-i][len(matrix)-1-j]
		}
		// rotate right to bottom
		for j := len(matrix)-1-i; j >= i; j-- {
			matrix[len(matrix)-1-i][len(matrix)-1-j] = matrix[j][len(matrix)-1-i]
		}
		// rotate top to right
		for j := len(matrix)-1-i; j >= i; j-- {
			matrix[j][len(matrix)-1-i] = tmp[j-i]
		}
	}
}
