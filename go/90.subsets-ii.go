/*
 * @lc app=leetcode id=90 lang=golang
 *
 * [90] Subsets II
 */

// @lc code=start
var subsets [][]int
func subsetsWithDup(nums []int) [][]int {
	subsets = nil
	sort.Ints(nums)
	for i := 0; i <= len(nums); i++ {
		findSubsetsWithNElements(nums, i, []int{})
	}
	return subsets
}

func findSubsetsWithNElements(nums []int, n int, subset []int) {
	if n == 0 {
		subsets = append(subsets, makeSliceCopy(subset))
	} else {
		if len(nums) < n {
			return
		}
		findSubsetsWithNElements(removeElementFromSlice(nums, 0), n-1, append(subset, nums[0]))
		for i := 1; i < len(nums); i++ {
			if nums[i] != nums[i-1] {
				findSubsetsWithNElements(nums[i+1:], n-1, append(subset, nums[i]))
			}
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

