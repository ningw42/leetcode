/*
 * @lc app=leetcode id=35 lang=golang
 *
 * [35] Search Insert Position
 */
func searchInsert(nums []int, target int) int {
	for i, v := range nums {
		if v < target {
			continue
		} else {
			return i
		}
	}

	return len(nums)
}

