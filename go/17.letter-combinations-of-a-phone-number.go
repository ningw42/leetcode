/*
 * [17] Letter Combinations of a Phone Number
 *
 * https://leetcode.com/problems/letter-combinations-of-a-phone-number/description/
 *
 * algorithms
 * Medium (38.52%)
 * Total Accepted:    295.6K
 * Total Submissions: 765.7K
 * Testcase Example:  '"23"'
 *
 * Given a string containing digits from 2-9 inclusive, return all possible
 * letter combinations that the number could represent.
 *
 * A mapping of digit to letters (just like on the telephone buttons) is given
 * below. Note that 1 does not map to any letters.
 *
 *
 *
 * Example:
 *
 *
 * Input: "23"
 * Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
 *
 *
 * Note:
 *
 * Although the above answer is in lexicographical order, your answer could be
 * in any order you want.
 *
 */
func letterCombinations(digits string) []string {
	// return letterCombinationsRecursion(digits)
	return letterCombinationsLoop(digits)
}

func letterCombinationsLoop(digits string) []string {
	m := map[string][]string{
		"2": []string{"a", "b", "c"},
		"3": []string{"d", "e", "f"},
		"4": []string{"g", "h", "i"},
		"5": []string{"j", "k", "l"},
		"6": []string{"m", "n", "o"},
		"7": []string{"p", "q", "r", "s"},
		"8": []string{"t", "u", "v"},
		"9": []string{"w", "x", "y", "z"},
	}

	var result []string

	for _, r := range digits {
		tmp := []string{}
		s := string(r)
		for _, suffix := range m[s] {
			if len(result) == 0 {
				tmp = m[s]
				break
			} else {
				for _, item := range result {
					tmp = append(tmp, item+suffix)
				}
			}
		}
		result = tmp
	}

	return result
}

func letterCombinationsRecursion(digits string) []string {
	m := map[string][]string{
		"2": []string{"a", "b", "c"},
		"3": []string{"d", "e", "f"},
		"4": []string{"g", "h", "i"},
		"5": []string{"j", "k", "l"},
		"6": []string{"m", "n", "o"},
		"7": []string{"p", "q", "r", "s"},
		"8": []string{"t", "u", "v"},
		"9": []string{"w", "x", "y", "z"},
	}

	if len(digits) == 0 {
		return []string{}
	}

	if len(digits) == 1 {
		return m[digits]
	}

	var result []string
	for _, item := range letterCombinations(digits[1:]) {
		for _, prefix := range m[digits[0:1]] {
			result = append(result, prefix+item)
		}
	}

	return result
}
