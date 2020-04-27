var count int
func totalNQueens(n int) int {
	count = 0
	solve(n, 0, [][]int{})
	return count
}

func solve(n, i int, queens [][]int) {
	if i == n {
		count++
	} else {
		for j := 0; j < n; j++ {
			if _, conflicts := calculateConflictPositions(queens, n)[j]; !conflicts {
				solve(n, i+1, append(queens, []int{i, j}))
			}
		}
	}
}

func calculateConflictPositions(queens [][]int, n int) map[int]bool {
	conflicts := make(map[int]bool)
	for i := len(queens)-1; i >= 0; i-- {
		y := queens[i][1]
		pos := []int{y, y-(len(queens)-i), y+(len(queens)-i)}
		for _, y := range pos {
			if y >= 0 && y < n {
				conflicts[y] = true
			}
		}
	}
	return conflicts
}