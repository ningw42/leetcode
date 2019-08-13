/*
 * @lc app=leetcode id=26 lang=golang
 *
 * [26] Remove Duplicates from Sorted Array
 */
func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	previous := nums[0]
	for i := 1; i < len(nums); {
		if previous == nums[i] {
			nums = append(nums[:i], nums[i+1:]...)
		} else {
			previous = nums[i]
			i++
		}
	}

	return len(nums)
}

