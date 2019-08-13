/*
 * @lc app=leetcode id=11 lang=golang
 *
 * [11] Container With Most Water
 */
func maxArea(heights []int) int {
	max := 0

	for i, left := range heights {
		for j := i + 1; j < len(heights); j++ {
			right := heights[j]
			length := j - i
			var height int
			if left < right {
				height = left
			} else {
				height = right
			}
			area := length * height
			if area > max {
				max = area
			}
		}
	}

	return max
}

