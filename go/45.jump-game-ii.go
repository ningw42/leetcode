/*
 * @lc app=leetcode id=45 lang=golang
 *
 * [45] Jump Game II
 */

// @lc code=start
func jump(nums []int) int {
	return dp(nums)
}

func dp(nums []int) int {
	steps := make([]int, len(nums))
	for i := len(nums) - 2; i >= 0; i-- {
		min := math.MaxInt32
		for j := nums[i]; j >= 1; j-- {
			if i+j >= len(nums) {
				min = 0
				break
			}
			if step := steps[i+j]; step < min {
				min = step
			}
		}
		steps[i] = min + 1
	}

	return steps[0]
}

func recursion(nums []int, current int) int {
	if current >= len(nums) - 1 {
		return 0
	}

	min := math.MaxInt32 
	for i := 1; i <= nums[current]; i++ {
		if count := recursion(nums, current+i); count < min {
			min = count
		}
	}

	return min + 1
}
// @lc code=end

