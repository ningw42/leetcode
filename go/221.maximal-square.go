func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}

	var result int
	var results [][]int
	for i := 0; i < len(matrix); i++ {
		results = append(results, make([]int, len(matrix[i])))
	}
	max := 0
	for i := 0; i < len(matrix); i++ {
		result = (int)(matrix[i][0] - '0')
		if result > max {
			max = result
		}
		results[i][0] = result
	}
	for i := 0; i < len(matrix[0]); i++ {
		result = (int)(matrix[0][i] - '0')
		if result > max {
			max = result
		}
		results[0][i] = result
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][j] == '1' {
				// min(results[i-1][j], results[i][j-1], results[i-1][j-1])
				r1 := results[i-1][j] + 1
				r2 := results[i][j-1] + 1
				r3 := results[i-1][j-1] + 1
				if r1 < r2 {
					result = r1
				} else {
					result = r2
				}
				if r3 < result {
					result = r3
				}
			} else {
				result = 0
			}
			results[i][j] = result
			if result > max {
				max = result
			}
		}
	}

	return max * max
}
