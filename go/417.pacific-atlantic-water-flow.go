/*
 * @lc app=leetcode id=417 lang=golang
 *
 * [417] Pacific Atlantic Water Flow
 */
func pacificAtlantic(matrix [][]int) [][]int {
	pacific := findPath(matrix, "pacific")
	atlantic := findPath(matrix, "atlantic")

	var ret [][]int
	for i, row := range matrix {
		for j, _ := range row {
			if pacific[i][j] && atlantic[i][j] {
				ret = append(ret, []int{i, j})
			}
		}
	}

	return ret
}

func findPath(matrix [][]int, ocean string) [][]bool {
	var ret [][]bool
	for _, row := range matrix {
		ret = append(ret, make([]bool, len(row)))
	}

	for i, row := range matrix {
		for j, _ := range row {
			if ocean == "pacific" {
				if i == 0 || j == 0 {
					DFS(matrix, ret, i, j)
				}
			} else {
				if i == len(matrix) - 1 || j == len(row) - 1 {
					DFS(matrix, ret, i, j)
				}
			}
		}
	}

	return ret
}

func DFS(matrix [][]int, visited [][]bool, i int, j int) {
	visited[i][j] = true
	coods := [][]int{[]int{i-1,j}, []int{i+1,j}, []int{i,j-1}, []int{i,j+1}}
	for _, cood := range coods {
		x, y := cood[0], cood[1]
		if 0 <= x && x < len(matrix) && 0 <= y && y < len(matrix[0]) {
			if matrix[x][y] >= matrix[i][j] && !visited[x][y] {
				DFS(matrix, visited, x, y)
			}
		}
	}
}

