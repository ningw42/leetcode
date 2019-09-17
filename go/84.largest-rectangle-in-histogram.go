/*
 * @lc app=leetcode id=84 lang=golang
 *
 * [84] Largest Rectangle in Histogram
 */
func largestRectangleArea(heights []int) int {
	return DP(heights)
}

func DP(heights []int) int {
	// TODO
}

func BruteForce(heights []int) int {
	max := 0
    for i := 0; i < len(heights); i++ {
		height := heights[i]
		for j := i; j >= 0; j-- {
			if heights[j] < height {
				height = heights[j]
			}
			area := height * (i - j + 1)
			if max < area {
				max = area
			}
		}
	}

	return max
}

