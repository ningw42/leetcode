/*
 * @lc app=leetcode id=169 lang=golang
 *
 * [169] Majority Element
 */

// @lc code=start
func majorityElement(nums []int) int {
	counts := make(map[int]int)
	var num, count int
	for i := 0; i < len(nums); i++ {
		num = nums[i]
		if old, exists := counts[num]; exists {
			count = old + 1
		} else {
			count = 1
		}
		if count > len(nums)/2 {
			break
		}
		counts[num] = count
	}
	return num
}

// @lc code=end

