/*
 * @lc app=leetcode id=119 lang=golang
 *
 * [119] Pascal's Triangle II
 */
func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	} else if rowIndex == 1 {
		return []int{1,1}
	}

	lastRow := []int{1,1}
	for i := 2; i <= rowIndex; i++ {
		currentRow := []int{1}
		for j := 1; j < i; j++ {
			currentRow = append(currentRow, lastRow[j-1] + lastRow[j])
		}
		currentRow = append(currentRow, 1)
		lastRow = currentRow
	}

	return lastRow
}

