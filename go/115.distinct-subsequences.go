/*
 * @lc app=leetcode id=115 lang=golang
 *
 * [115] Distinct Subsequences
 */
func numDistinct(s string, t string) int {
	return DP(s, t)
}

// consider from tail to head
// the number of T in S is
//   the number of T in S[1:] if T[0] != S[0]
//   the number of T in S[1:0] + the number of T[1:0] in S[1:0] if T[0] == S[0]

// example:
// ""	g	ag	bag	gbag	bgbag	abgbag	babgbag	bbabgbag
// ""	1	1	1	1	1		1		1		1		1
// g	0	1	1	1	2		2		2		2		2
// ag	0	0	1	1	1		1		3		3		3
// bag	0	0	0	1	1		2		2		5		8
func DP(s, t string) int {
	var results [][]int
	for i := 0; i <= len(t); i++ {
		results = append(results, make([]int, len(s)+1))
	}
	for i := 0; i <= len(s); i++ {
		results[0][i] = 1
	}

	var result, I, J int
	for i := 1; i <= len(t); i++ {
		for j := i; j <= len(s); j++ {
			I = len(t) - i
			J = len(s) - j
			if s[J] == t[I] {
				result = results[i][j-1] + results[i-1][j-1]
			} else {
				result = results[i][j-1]
			}
			results[i][j] = result
		}
	}

	return results[len(t)][len(s)]
}

// too much repeative calculation just like recursive Fib
func Recursion(s, t string) int {
	if len(t) == 0 {
		return 1
	}

	result := 0
	for j := 0; j < len(s); j++ {
		if t[0] == s[j] {
			result += numDistinct(s[j+1:], t[1:])
		}
	}

	return result
}

