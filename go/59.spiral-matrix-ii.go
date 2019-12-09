/*
 * @lc app=leetcode id=59 lang=golang
 *
 * [59] Spiral Matrix II
 */

// @lc code=start

const GoRight int = 0
const GoDown int = 1
const GoLeft int = 2
const GoUp int = 3
func generateMatrix(n int) [][]int {
	var matrix [][]int
	for i := 0; i < n; i++ {
		matrix = append(matrix, make([]int, n))
	}

	current := 1
	square := n * n
	direction := 0
	x, y := 0, 0
	for current <= square {
		switch direction {
		case GoRight:
			for y < n && matrix[x][y] == 0 {
				matrix[x][y] = current
				current++
				y++
			}
			y--
			x++
		case GoDown:
			for x < n && matrix[x][y] == 0 {
				matrix[x][y] = current
				current++
				x++
			}
			x--
			y--
		case GoLeft:
			for y >= 0 && matrix[x][y] == 0 {
				matrix[x][y] = current
				current++
				y--
			}
			y++
			x--
		case GoUp:
			for x >= 0 && matrix[x][y] == 0 {
				matrix[x][y] = current
				current++
				x--
			}
			x++
			y++
		}
		direction = (direction+1)%4
	}

	return matrix
}
// @lc code=end

