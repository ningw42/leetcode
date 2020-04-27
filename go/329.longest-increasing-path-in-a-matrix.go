func longestIncreasingPath(matrix [][]int) int {
	var result [][]int
	for _, row := range matrix {
		result = append(result, make([]int, len(row)))
	}

	max := 0
	for i, row := range matrix {
		for j, _ := range row {
			if dist := find(matrix, result, i, j); dist > max {
				max = dist
			}
		}
	}

	return max
}

func find(matrix, result [][]int, i, j int) int {
	dist := result[i][j]
	if dist != 0 {
		return dist
	}

	coods := [][]int{
		[]int{-1, 0},
		[]int{1, 0},
		[]int{0, -1},
		[]int{0, 1},
	}
	for _, cood := range coods {
		x := i + cood[0]
		y := j + cood[1]
		if 0 <= x && x < len(matrix) && 0 <= y && y < len(matrix[i]) && matrix[x][y] > matrix[i][j] {
			if d := find(matrix, result, x, y); d >= dist {
				dist = d
			}
		}
	}

	result[i][j] = dist + 1

	return result[i][j]
}
