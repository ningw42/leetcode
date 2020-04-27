
// S(i, j) = min(S(i, k) + S(k+1, j) + MAX(i, k) * MAX(k+1, j)) for every k with i < k < j
// MAX(i, j) = max(arr[i]...arr[j])
func mctFromLeafValues(arr []int) int {
	maxs := make([][]int, len(arr))
	for i := 0; i < len(arr); i++ {
		maxs[i] = make([]int, len(arr))
	}
	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			if j == i {
				maxs[i][j] = arr[i]
			} else {
				maxs[i][j] = max(maxs[i][j-1], arr[j])
			}
		}
	}

	mcts := make([][]int, len(arr))
	for i := 0; i < len(arr); i++ {
		mcts[i] = make([]int, len(arr))
	}
	for step := 1; step < len(arr); step++ {
		for i := 0; i < len(arr); i++ {
			j := i + step
			if j >= len(arr) {
				continue
			}
			result := math.MaxInt32
			for k := i; k < j; k++ {
				result = min(
					result,
					mcts[i][k] + mcts[k+1][j] + maxs[i][k] * maxs[k+1][j],
				)
			}
			mcts[i][j] = result
		}
	}

	return mcts[0][len(arr)-1]
}

func max(a, b int) int {
	if a > b {
		return a 
	} else {
		return b
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}