/*
 * @lc app=leetcode id=18 lang=golang
 *
 * [18] 4Sum
 */

// @lc code=start
func fourSum(nums []int, target int) [][]int {
	var result [][]int
	sort.Ints(nums)

	for i := 0; i < len(nums) - 3; i++ {
		if i == 0 || nums[i-1] != nums[i] {
			if part := threeSum(nums[i+1:], target - nums[i]); len(part) != 0 {
				for _, three := range part {
					result = append(result, append([]int{nums[i]}, three...))
				}
			}
		}
	}

	return result
}

func threeSum(nums []int, target int) [][]int {
	var result [][]int

	for i := 0; i < len(nums) - 2; i++ {
		// the first element of result fixed
		// skip all elements with the same value
		if i == 0 || nums[i-1] != nums[i] {
			t := target - nums[i]
			j := i + 1
			k := len(nums) - 1
			for j < k {
				sum := nums[j] + nums[k]
				if sum > t {
					// skip elements with same value
					for j < k && nums[k-1] == nums[k] {k--}
					k--
				} else if sum < t {
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

