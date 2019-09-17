/*
 * @lc app=leetcode id=85 lang=golang
 *
 * [85] Maximal Rectangle
 */
func maximalRectangle(matrix [][]byte) int {
	return DynamicProgramming(matrix)
}


// reference: https://leetcode.com/problems/maximal-rectangle/discuss/29054/Share-my-DP-solution
func DynamicProgramming(matrix [][]byte) int {
	max := 0
	lenX := len(matrix)
	if lenX == 0 {
		return 0
	}
	lenY := len(matrix[0])
	var height, left, right []int
	for i := 0; i < lenY; i++ {
		left = append(left, 0)
		right = append(right, lenY)
		height = append(height, 0)
	}

	for i := 0; i < lenX; i++ {
		// find max height(consecutive '1's upwards)
		for j := 0; j < lenY; j++ {
			if matrix[i][j] == '1' {
				height[j] = height[j] + 1
			} else {
				height[j] = 0
			}
		}
		maxLeft := 0
		for j := 0; j < lenY; j++ {
			if matrix[i][j] == '1' {
				if maxLeft > left[j] {
					left[j] = maxLeft
				}
			} else {
				left[j] = 0
				maxLeft = j + 1
			}
		}
		minRight := lenY
		for j := lenY - 1; j >= 0; j-- {
			if matrix[i][j] == '1' {
				if minRight < right[j] {
					right[j] = minRight
				}
			} else {
				right[j] = lenY
				minRight = j
			}
		}

		for j := 0; j < lenY; j++ {
			area := height[j] * (right[j] - left[j])
			if area > max {
				max = area
			}
		}
	}

	return max
}

func BruteForce(matrix [][]byte) int {
	max := 0
    for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			for m := i; m < len(matrix); m++ {
				for n := j; n < len(matrix[0]); n++ {
					if r := getRectangle(matrix, i, j, m, n); max < r {
						max = r
					}
				}
			}
		}
	}

	return max
}

func getRectangle(matrix [][]byte, i, j, m, n int) int {
	for a := i; a <= m; a++ {
		for b := j; b <= n; b++ {
			if matrix[a][b] == '0' {
				return 0
			}
		}
	}
	return (m-i+1) * (n-j+1)
}

