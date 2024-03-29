/*
 * @lc app=leetcode id=946 lang=golang
 *
 * [946] Validate Stack Sequences
 *
 * https://leetcode.com/problems/validate-stack-sequences/description/
 *
 * algorithms
 * Medium (57.93%)
 * Total Accepted:    7.1K
 * Total Submissions: 12.3K
 * Testcase Example:  '[1,2,3,4,5]\n[4,5,3,2,1]'
 *
 * Given two sequences pushed and popped with distinct values, return true if
 * and only if this could have been the result of a sequence of push and pop
 * operations on an initially empty stack.
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: pushed = [1,2,3,4,5], popped = [4,5,3,2,1]
 * Output: true
 * Explanation: We might do the following sequence:
 * push(1), push(2), push(3), push(4), pop() -> 4,
 * push(5), pop() -> 5, pop() -> 3, pop() -> 2, pop() -> 1
 *
 *
 *
 * Example 2:
 *
 *
 * Input: pushed = [1,2,3,4,5], popped = [4,3,5,1,2]
 * Output: false
 * Explanation: 1 cannot be popped before 2.
 *
 *
 *
 *
 * Note:
 *
 *
 * 0 <= pushed.length == popped.length <= 1000
 * 0 <= pushed[i], popped[i] < 1000
 * pushed is a permutation of popped.
 * pushed and popped have distinct values.
 *
 *
 *
 */
func validateStackSequences(pushed []int, popped []int) bool {
	var s []int
	popIndex := 0
	
	// imitating the push and pop process
	// if the current top element == the next poped element, pop and continue the step on top of the stack
	// until the current top element != the next poped element, push next element and repeat the entire procedure
	for _, nextPushedElement := range pushed {
		s = append(s, nextPushedElement)
		for len(s) != 0 && popIndex < len(popped) && s[len(s)-1] == popped[popIndex] {
			s = s[0 : len(s)-1]
			popIndex++
		}
	}

	// if the given sequence is possible, the popIndex should be equal to length of the pop sequence
	return popIndex == len(popped)
}
