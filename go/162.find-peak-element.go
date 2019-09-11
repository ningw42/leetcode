/*
 * @lc app=leetcode id=162 lang=golang
 *
 * [162] Find Peak Element
 */
func findPeakElement(nums []int) int {
	length := len(nums)
	for i := 0; i < length; i++ {
		left := get(nums, length, i-1)
		mid := get(nums, length, i)
		right := get(nums, length, i+1)
		if left < mid && mid > right {
			return i
		}
	}
	return 0
}

func get(nums []int, length, index int) int {
	if index < 0 || index >= length {
		return math.MinInt32
	} else {
		return nums[index]
	}
}

