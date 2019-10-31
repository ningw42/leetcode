/*
 * @lc app=leetcode id=53 lang=golang
 *
 * [53] Maximum Subarray
 */

// @lc code=start
func maxSubArray(nums []int) int {
	var result int
	sums := make([]int, len(nums))
	sums[0] = nums[0]
	max := sums[0]
	for i := 1; i < len(nums); i++ {
		if s := sums[i-1] + nums[i]; s > nums[i] {
			result = s
		} else {
			result = nums[i]
		}
		sums[i] = result
		if result > max {
			max = result
		}
	}

	return max
}

// @lc code=end

