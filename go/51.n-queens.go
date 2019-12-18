/*
 * @lc app=leetcode id=51 lang=golang
 *
 * [51] N-Queens
 */

// @lc code=start
var solutions [][][]int
func solveNQueens(n int) [][]string {
	solutions = nil
	solve(n, 0, [][]int{})

	var ret [][]string
	for _, solution := range solutions {
		ret = append(ret, makeSolution(solution))
	}
	return ret
}

func makeSolution(solution [][]int) []string {
	var s [][]byte
	for i := 0; i < len(solution); i++ {
		var row []byte
		for j := 0; j < len(solution); j++ {
			row = append(row, '.')
		}
		s = append(s, row)
	}
	for _, queen := range solution {
		x, y := queen[0], queen[1]
		s[x][y] = 'Q'
	}

	var ret []string
	for _, row := range s {
		ret = append(ret, string(row))
	}
	return ret
}

func solve(n, i int, queens [][]int) {
	if i == n {
		solutions = append(solutions, makeCopy(queens))
	} else {
		for j := 0; j < n; j++ {
			if _, conflicts := calculateConflictPositions(queens, n)[j]; !conflicts {
				solve(n, i+1, append(queens, []int{i, j}))
			}
		}
	}
}

func makeCopy(s [][]int) [][]int {
	var ret [][]int 
	for _, ss := range s {
		ret = append(ret, ss)
	}
	return ret
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
// @lc code=end

