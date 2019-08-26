/*
 * @lc app=leetcode id=118 lang=golang
 *
 * [118] Pascal's Triangle
 */
func generate(numRows int) [][]int {
	if numRows == 0 {
		return [][]int{}
	} else if numRows == 1 {
		return [][]int{[]int{1}}
	} else if numRows == 2 {
		return [][]int{[]int{1}, []int{1, 1}}
	}

	result := [][]int{[]int{1}, []int{1, 1}}
	for i := 3; i <= numRows; i++ {
		row := []int{1}
		for j := 1; j < i - 1; j++ {
			row = append(row, result[i-2][j-1] + result[i-2][j])
		}
		row = append(row, 1)
		result = append(result, row)
	}

	return result
}

