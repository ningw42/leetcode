/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 */

// @lc code=start
func threeSum(nums []int) [][]int {
	var result [][]int
	sort.Ints(nums)

	for i := 0; i < len(nums) - 2; i++ {
		// the first element of result fixed
		// skip all elements with the same value
		if i == 0 || nums[i-1] != nums[i] {
			target := -nums[i]
			j := i + 1
			k := len(nums) - 1
			for j < k {
				sum := nums[j] + nums[k]
				if sum > target {
					// skip elements with same value
					for j < k && nums[k-1] == nums[k] {k--}
					k--
				} else if sum < target {
					// skip elements with same value
					for j < k && nums[j+1] == nums[j] {j++}
					j++
				} else {
					result = append(result, []int{nums[i], nums[j], nums[k]})
					// skip elements with same value
					for j < k && nums[k-1] == nums[k] {k--}
					k--
					// skip elements with same value
					for j < k && nums[j+1] == nums[j] {j++}
					j++
				}
			}
		}
	}

	return result
}

// @lc code=end

