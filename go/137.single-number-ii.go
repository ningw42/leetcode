/*
 * @lc app=leetcode id=137 lang=golang
 *
 * [137] Single Number II
 */

// @lc code=start
func singleNumber(nums []int) int {
	return withExtraMemory(nums)    
}

func withExtraMemory(nums []int) int {
	dist := make(map[int]int)
	for _, num := range nums {
		if _, exists := dist[num]; !exists {
			dist[num] = 1
		} else {
			dist[num]++
		}
	}

	for num, c := range dist {
		if c == 1 {
			return num
		}
	}
	return 0
}

func withoutExtraMemory(nums []int) int {
	// TODO https://leetcode.com/problems/single-number-ii/discuss/43295/Detailed-explanation-and-generalization-of-the-bitwise-operation-method-for-single-numbers
	return 0
}
// @lc code=end

