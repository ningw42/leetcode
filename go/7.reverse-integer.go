import (
	"strconv"
)

/*
 * [7] Reverse Integer
 *
 * https://leetcode.com/problems/reverse-integer/description/
 *
 * algorithms
 * Easy (24.45%)
 * Total Accepted:    498.2K
 * Total Submissions: 2M
 * Testcase Example:  '123'
 *
 * Given a 32-bit signed integer, reverse digits of an integer.
 *
 * Example 1:
 *
 *
 * Input: 123
 * Output: 321
 *
 *
 * Example 2:
 *
 *
 * Input: -123
 * Output: -321
 *
 *
 * Example 3:
 *
 *
 * Input: 120
 * Output: 21
 *
 *
 * Note:
 * Assume we are dealing with an environment which could only store integers
 * within the 32-bit signed integer range: [−231,  231 − 1]. For the purpose of
 * this problem, assume that your function returns 0 when the reversed integer
 * overflows.
 *
 */
func reverse(x int) int {
	if x == 0 {
		return 0
	}

	sign := x >> 31 & 1
	var unsigned uint64
	if sign == 1 {
		unsigned = uint64(-x)
	} else {
		unsigned = uint64(x)
	}

	b := []byte(strconv.FormatUint(unsigned, 10))
	r := []byte{}
	if sign == 1 {
		r = append(r, '-')
	}

	for i := len(b) - 1; i >= 0; i-- {
		r = append(r, b[i])
	}

	reversed, _ := strconv.ParseInt(string(r), 10, 64)

	// overflow
	if reversed < -0x7FFFFFFF || reversed > 0x7FFFFFFE {
		return 0
	}

	return int(reversed)
}
