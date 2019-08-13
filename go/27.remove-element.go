/*
 * @lc app=leetcode id=27 lang=golang
 *
 * [27] Remove Element
 */
func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	
	for i := 0; i < len(nums) - 1; {
		v := nums[i]
		if v == val {
			nums = append(nums[:i], nums[i+1:]...)
		} else {
			i++
		}
	}

	if nums[len(nums) - 1] == val {
		nums = nums[:len(nums) - 1]
	}

	return len(nums)
}

