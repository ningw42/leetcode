func orderOfLargestPlusSign(N int, mines [][]int) int {
	mat := buildMatrix(N, mines)
	max := math.MinInt32
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			order := findOrder(mat, i, j)
			// fmt.Println(i, j, order)
			if order > max {
				max = order
			}
		}
	}

	return max
}

func findOrder(mat [][]bool, i, j int) int {
	if mat[i][j] {
		return 0
	}

	N := len(mat)
	diff := 1
	var x, y int
loop:
	for {
		coods := [][]int{
			[]int{i + diff, j},
			[]int{i - diff, j},
			[]int{i, j + diff},
			[]int{i, j - diff},
		}
		for _, cood := range coods {
			x = cood[0]
			y = cood[1]
			if 0 <= x && x < N && 0 <= y && y < N {
				if mat[x][y] {
					break loop
				}
			} else {
				break loop
			}
		}
		diff++
	}

	return diff
}

func buildMatrix(N int, mines [][]int) [][]bool {
	var mat [][]bool
	for i := 0; i < N; i++ {
		mat = append(mat, make([]bool, N))
	}
	for _, cood := range mines {
		mat[cood[0]][cood[1]] = true
	}

	return mat
}
