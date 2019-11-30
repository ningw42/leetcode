/*
 * @lc app=leetcode id=46 lang=golang
 *
 * [46] Permutations
 */

// @lc code=start

var result [][]int

func permute(nums []int) [][]int {
	result = [][]int{}
	findPermute([]int{}, nums)

	return result
}

func findPermute(permute []int, nums []int) {
	if len(nums) == 0 {
		result = append(result, makeSliceCopy(permute))
	} else {
		for i, num := range nums {
			findPermute(append(permute, num), removeElementFromSlice(nums, i))
		}
	}
}

func makeSliceCopy(s []int) []int {
	copy := make([]int, len(s))
	for i, e := range s {
		copy[i] = e
	}
	return copy
}

func removeElementFromSlice(s []int, i int) []int {
	var copy []int
	return append(append(copy, s[0:i]...), s[i+1:]...)
}
// @lc code=end

