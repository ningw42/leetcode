import "sort"

/*
 * [22] Generate Parentheses
 *
 * https://leetcode.com/problems/generate-parentheses/description/
 *
 * algorithms
 * Medium (50.28%)
 * Total Accepted:    253.8K
 * Total Submissions: 502K
 * Testcase Example:  '3'
 *
 *
 * Given n pairs of parentheses, write a function to generate all combinations
 * of well-formed parentheses.
 *
 *
 *
 * For example, given n = 3, a solution set is:
 *
 *
 * [
 * ⁠ "((()))",
 * ⁠ "(()())",
 * ⁠ "(())()",
 * ⁠ "()(())",
 * ⁠ "()()()"
 * ]
 *
 */
func generateParenthesis(n int) []string {
	if n == 1 {
		return []string{"()"}
	} else if n == 2 {
		return []string{"()()", "(())"}
	} else {
		m := make(map[string]bool)
		result := []string{}
		for _, v := range generateParenthesis(n - 1) {
			newStrings := insert(v)
			for _, s := range newStrings {
				m[s] = true
			}
		}

		for k, _ := range m {
			result = append(result, k)
		}
		sort.Sort(sort.StringSlice(result))

		return result
	}
}

func insert(s string) []string {
	result := []string{"()" + s, s + "()"}

	for p := range s {

		b := []byte(s)
		if p < len(s)-1 {
			result = append(result, string(b[0:p+1])+"()"+string(b[p+1:]))
		}
	}

	return result
}
