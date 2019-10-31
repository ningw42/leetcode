/*
 * @lc app=leetcode id=72 lang=golang
 *
 * [72] Edit Distance
 */

// @lc code=start
func minDistance(word1 string, word2 string) int {
	var mat [][]int
	for i := 0; i <= len(word1); i++ {
		mat = append(mat, make([]int, len(word2)+1))
	}
	for i := 0; i <= len(word1); i++ {
		mat[i][0] = i
	}
	for i := 0; i <= len(word2); i++ {
		mat[0][i] = i
	}

	var result int
	for i := 0; i < len(word1); i++ {
		for j := 0; j < len(word2); j++ {
			if word1[i] == word2[j] {
				result = mat[i][j]
			} else {
				result = mat[i][j] + 1
				if del := mat[i][j+1] + 1; del < result {
					result = del
				}
				if ins := mat[i+1][j] + 1; ins < result {
					result = ins
				}
			}
			mat[i+1][j+1] = result
		}
	}

	return mat[len(word1)][len(word2)]
}

// @lc code=end

