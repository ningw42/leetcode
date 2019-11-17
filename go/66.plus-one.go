/*
 * @lc app=leetcode id=66 lang=golang
 *
 * [66] Plus One
 */

// @lc code=start
func plusOne(digits []int) []int {
	var carry, digit int
	carry = 1
	for i := len(digits) - 1; i >= 0; i-- {
		digit = digits[i] + carry
		if digit >= 10 {
			digit = digit - 10
			carry = 1
		} else {
			carry = 0
		}
		digits[i] = digit
	}
	if carry == 1 {
		digits = append([]int{1}, digits...)
	}

	return digits
}
// @lc code=end

